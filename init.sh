#!/usr/bin/env bash
export LC_CTYPE=C
export LANG=C

REPO_NAME="gank-global"
REPO_HOST="bitbucket.org\/gank-global"

function display_help {
    echo "Usage: $0 [option...] create { project_name } { path }" >&2
    echo
    echo "   -h, --help                 Display help information"
    echo
}

function create_project {
    if [ -z $2 ]; then
        display_help
        exit 2
    fi
    echo "Creating new project..."
    PROJ_PATH=$2/$1
    mkdir -p $PROJ_PATH

    cp -R ./src/ProjectTemplate/ $PROJ_PATH
    rm -rf $PROJ_PATH/vendor

    # grep -R "bitbucket.org/$REPO_NAME/ProjectTemplate" -l --null $PROJ_PATH | xargs -0 sed -i '' -e "s/ProjectTemplate/$1/g"
    find $PROJ_PATH -type f -exec sed -i '' -e "s/\${REPO_HOST}/$REPO_HOST/g" \
        -e "s/\${FILENAME}/sample/g" \
        -e "s/\${PROJ_NAME}/$1/g" \
        -e "s/\${DASHED_NAME}/sample/g" \
        -e "s/\${SERVICE_NAME}/Sample/g" \
        -e "s/\${CAMELIZED_NAME}/sample/g" \
            {} \;

    echo "Project $1 created."
    echo "Path $PROJ_PATH"
}

function gen_service {
    SERVICE_NAME=$1
    INIT_PATH=$(echo $2 | sed -r 's/\/+/\//g' | sed -r 's/\/$//g')
    cd $INIT_PATH
    PROJ_NAME=$(echo "${PWD##*/}")
    echo "$PROJ_NAME"
    cd -

    #FILENAME: order order_item ...
    FILENAME=$(echo $1 | sed 's/\(.\)\([A-Z]\)/\1_\2/g' | tr '[:upper:]' '[:lower:]')
    DASHED_NAME=$(echo $1 | sed 's/\(.\)\([A-Z]\)/\1-\2/g' | tr '[:upper:]' '[:lower:]')
    CAMELIZED_NAME="$(tr '[:upper:]' '[:lower:]' <<< ${SERVICE_NAME:0:1})${SERVICE_NAME:1}"

    echo "Generate service $DASHED_NAME ..."
    DIRECTORIES=( "model" "repository" "repository/postgresql" "service/sample" "service/sample/dto" "handler" )

    for SRC in "${DIRECTORIES[@]}"
    do
        DST=$(echo $SRC | sed "s/sample/$FILENAME/g")
        DST_PATH=$INIT_PATH/$DST
        mkdir -p $DST_PATH
        for i in ./src/ProjectTemplate/$SRC/sample*
        do
            FILE_MAPPING=$(echo $i \
                | sed "s/sample/$FILENAME/g" \
                | sed "s/\.\/src\/ProjectTemplate//g")

            FILE_PATH="$INIT_PATH/$FILE_MAPPING"
            echo "Creating file $FILE_PATH ..."
            cp "$i" "$FILE_PATH"
        done

        FILE_MASK="${FILENAME}*.go"

        find $DST_PATH -type f -exec sed -i '' -e "s/\${REPO_HOST}/$REPO_HOST/g" \
            -e "s/\${FILENAME}/$FILENAME/g" \
            -e "s/\${PROJ_NAME}/$PROJ_NAME/g" \
            -e "s/\${DASHED_NAME}/$DASHED_NAME/g" \
            -e "s/\${SERVICE_NAME}/$SERVICE_NAME/g" \
            -e "s/\${CAMELIZED_NAME}/$CAMELIZED_NAME/g" \
             {} \;
    done

    echo "Generating router ..."
    echo "Updating $INIT_PATH/cmd/http/http.go"
    # modify http.go 
    REPO="${CAMELIZED_NAME}Repo := postgresql.New${SERVICE_NAME}Repository(db)"
    SERVICE="${CAMELIZED_NAME}Service := ${FILENAME}.New${SERVICE_NAME}Service(${CAMELIZED_NAME}Repo, internalValidator)"

    HANDLERS=""
    HANDLER_ROUTERS=""
    for i in "List" "Get" "Post" "Put" "Patch" "Delete"
    do
        HANDLERS+="\n\t\t${SERVICE_NAME}${i}Handler: internalMiddleware.MakeHandler(handler.New${SERVICE_NAME}${i}Handler(${CAMELIZED_NAME}Service).Handle),"
        HANDLER_ROUTERS+="\t${SERVICE_NAME}${i}Handler http.HandlerFunc\n"
    done

    IMPORT_SERVICE="\"bitbucket.org/$REPO_NAME/$PROJ_NAME/service/$FILENAME\""

    echo $INIT_PATH/cmd/http/http.go
    awk -v x="\n\t$REPO\n\t$SERVICE\n\n\thandlerFuncs := internalRouter.HandlerFuncs{"  '/handlerFuncs[[:space:]]/{f=1} !f{print} f{print x; f=0}' $INIT_PATH/cmd/http/http.go > $INIT_PATH/cmd/http/.http.go
    awk -v x="$HANDLERS"  '/HandlerFuncs+\{/{f=1} !f{print} !/}/&&f{print} /}/&&f{print x"\n\t}";f=0}' $INIT_PATH/cmd/http/.http.go > $INIT_PATH/cmd/http/http.go
    awk -v x="$IMPORT_SERVICE"  '/import/{f=1} !f{print} !/)/&&f{print} /)/&&f{print "\t"x"\n)";f=0}' $INIT_PATH/cmd/http/http.go > $INIT_PATH/cmd/http/.http.go
    cat $INIT_PATH/cmd/http/.http.go > $INIT_PATH/cmd/http/http.go
    rm $INIT_PATH/cmd/http/.http.go

    #modify router.go
    echo "Updating $INIT_PATH/router/router.go"
    awk -v x="\n$HANDLER_ROUTERS"  '/type[[:space:]]HandlerFuncs/{f=1} !f{print} !/}/&&f{print} /}/&&f{print x"\n}";f=0}' $INIT_PATH/router/router.go > $INIT_PATH/router/.router.go

    ROUTES=(
        "\t\tr.Route(\"/v1/${PROJ_NAME}/${DASHED_NAME}s\", func(r chi.Router) {"
            "\t\t\tr.Get(\"/\", handlerFuncs.${SERVICE_NAME}ListHandler)"
            "\t\t\tr.Get(\"/{id}\", handlerFuncs.${SERVICE_NAME}GetHandler)"
            "\t\t\tr.Put(\"/{id}\", handlerFuncs.${SERVICE_NAME}PutHandler)"
            "\t\t\tr.Patch(\"/{id}\", handlerFuncs.${SERVICE_NAME}PatchHandler)"
            "\t\t\tr.Delete(\"/{id}\", handlerFuncs.${SERVICE_NAME}DeleteHandler)"
        "\t\t})"
    )

    ROUTE_STRING=""
    for i in "${ROUTES[@]}"
    do
        ROUTE_STRING+="${i}\n"
    done

    awk -v r="$ROUTE_STRING" \
    'BEGIN {f = -1; v = 0}
    /api.Group/{f=1; v=0;}
    f == 1 {x = gsub(/\(/,"\("); y = gsub(/\)/,"\)"); v = v + x - y; }
    v != 0 && f == 1 {print}
    v == 0 && f == 1 {print r; f = -1}
    f == -1 {print}
    ' $INIT_PATH/router/.router.go > $INIT_PATH/router/router.go
    rm $INIT_PATH/router/.router.go

    echo "Service generated."
}

while :
do
    case "$1" in
        -h | --help)
            display_help
            exit 0
            ;;
        create)
            create_project $2 $3
            exit 0
            ;;
        gen)
            case "$2" in
                service)
                    # ./init.sh gen service User /path
                    gen_$2 ${@:3}
                    exit
                    ;;
                *)
                    break
                    ;;
            esac
            exit 0
            ;;
        *)
            break
            ;;
    esac
done