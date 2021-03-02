${SERVICE_NAME}s service
===========

## Requirements

### Install docker, docker-compose

    brew update
    brew install docker docker-compose

### Build & Run

    docker-compose build
    docker-compose up
    docker-compose down (required to down all services and networks)
    
## How to create handler
### Register route
    File: router/router.go
###
    ...
    type HandlerFuncs struct {
        ${SERVICE_NAME}Handler        http.HandlerFunc
    }
    
    func Create(handlerFuncs HandlerFuncs) http.Handler {
        api := chi.NewRouter()
    
        api.Group(func(r chi.Router) {
            r.Route("${CAMELIZED_NAME}", func(r chi.Router) {
                r.Post("/${FILENAME}", ${SERVICE_NAME}Handler)
            })
        })
        return api
    }
### Register HTTP
    File: cmd/http/http.go
###
    ...
    func main() {
        ...
        ${CAMELIZED_NAME}Repo := postgresql.New${SERVICE_NAME}Repository(db)
        ${CAMELIZED_NAME}Service := ${FILENAME}.New${SERVICE_NAME}Service(${CAMELIZED_NAME}Repo)
        handlerFuncs := internalRouter.HandlerFuncs{
            ${SERVICE_NAME}Handler:    internalMiddleware.MakeHandler(handler.New${SERVICE_NAME}Handler(${CAMELIZED_NAME}Service).Handle),
        }
        ...
    }
### Make Handler
    Directory: handler/${FILENAME}_get.go
###
    package handler
    
    type ${SERVICE_NAME}Handler struct {
    	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service
    }
    
    func New${SERVICE_NAME}Handler(${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service) ${SERVICE_NAME}Handler {
    	return ${SERVICE_NAME}Handler{
    	    ${CAMELIZED_NAME}Service: ${CAMELIZED_NAME}Service,
    	}
    }
    
    // @Tags ${SERVICE_NAME}s
    // @Router /${FILENAME}s [GET]
    // @ID ${SERVICE_NAME}GetHandler
    // @Summary Get ${CAMELIZED_NAME}
    // @Description Get ${CAMELIZED_NAME}
    // @Param request query filter.${SERVICE_NAME}GetListFilter true "${SERVICE_NAME} get list filter"
    // @Success 200 "OK" {object} model.${SERVICE_NAME}
    // @Failure 400 {object} response.Error "Bad Request"
    // @Failure 500 {object} response.Error "Internal Server Error"
    func (h ${SERVICE_NAME}Handler) Handle(r *http.Request) response.ApiResponse {
        output, err := h.${CAMELIZED_NAME}Service.GetAll(r.Context())
    	if err != nil {
    		return response.ConvertServiceError(err)
    	}
    	return response.Accepted(output)
    }

### Make Service
    File: services/${FILENAME}/${FILENAME}.go
###
    package ${FILENAME}
    
    type ${SERVICE_NAME}Service interface {
    	GetAll(ctx context.Context) (*model.${SERVICE_NAME}, error)
    }
    
    type ${CAMELIZED_NAME}Service struct {
    	${CAMELIZED_NAME}Repo repository.${SERVICE_NAME}Repository
    }
    
    func New${SERVICE_NAME}Service(
    	${CAMELIZED_NAME}Repo repository.${SERVICE_NAME}Repository,
    ) ${SERVICE_NAME}Service {
    	return ${CAMELIZED_NAME}Service{${CAMELIZED_NAME}Repo: ${CAMELIZED_NAME}Repo}
    }
    
    func (s ${CAMELIZED_NAME}Service) GetAll(ctx context.Context) (*model.${SERVICE_NAME}, error) {
    	// Todo here
    }
    

