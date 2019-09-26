package produtos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//owner do handler
type Controller struct {
	storage MySQLStorage
}

//constructor do nosso controller
func NewPproduto(stg MySQLStorage) *Controller {
	return &Controller{
		storage: stg,
	}
}

//endpoint que busca os veiculos
func (ctrl *Controller) Get(c *gin.Context) {
	produtos, err := ctrl.storage.GetProduto()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, produtos)
}

//endpoint que cria novos veiculos
func (ctrl *Controller) Create(c *gin.Context) {
	var v produtos
	//transforma a request em um objeto do tipo Veiculo
	if err := c.ShouldBindJSON(&v); err != nil {
		c.AbortWithStatusJSONAbortWithError(http.StatusBadRequest, err)
		return
	}
	//salva os dados no banco
	err := ctrl.storage.CreateProduto(v.Nome, v.Modelo, v.Marca, v.Peco)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

//atualiza veiculos
func (ctrl *Controller) Update(c *gin.Context) {
	var v Produto
	//transforma a request em um objeto do tipo Veiculo
	if err := c.ShouldBindJSON(&v); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	//salva os dados no banco
	err := ctrl.storage.UpdateProduto(v.ID, &v)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

//apaga um veiculo
func (ctrl *Controller) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//declara a variavel e ao mesmo tempo verifica se é diferente de nil
	if err := ctrl.storage.Delete(id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
