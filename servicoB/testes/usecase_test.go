package testes

import (
	"context"
	"errors"
	"modapilab1/internal/controller"
	"modapilab1/internal/domain/dto"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) FindData(ctx context.Context, zipcode string) (*dto.ResultOutpurDto, error) {
	args := mock.Called(ctx, zipcode)
	return args.Get(0).(*dto.ResultOutpurDto), args.Error(1)
}

func TestInvalidZipCode(t *testing.T) {
	//mock
	mockUseCase := new(MockUseCase)

	input := "89037501"
	//input := "69880000"
	expectedOutput := &dto.ResultOutpurDto{
		Temp_C: 22.7,
		Temp_F: 72.6,
		Temp_K: 295.85,
	}

	//Configura o mock
	mockUseCase.On("FindData", input).Return(expectedOutput, nil)

	//Controller
	ctrl := &controller.FindDataController{
		FindDataUseCase: mockUseCase,
	}

	req, err := http.NewRequest("GET", "/?cep="+input, nil)
	assert.NoError(t, err)

	// Criar um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()

	// Chamar o controlador
	handler := http.HandlerFunc(ctrl.FindData)
	handler.ServeHTTP(rr, req)

	// Verificar o código de status e a resposta
	assert.Equal(t, http.StatusOK, rr.Code)
	//expected := `{"Temp_C":22.7,"Temp_F":72.6,"Temp_K":295.85}`
	//assert.JSONEq(t, expected, rr.Body.String())

	// Verificar se o método FindData foi chamado com o parâmetro correto
	//mockUseCase.AssertCalled(t, "FindData", input)
}

func TestValidZipCode200(t *testing.T) {
	input := "89037501"

	//Controller
	ctrl := &controller.FindDataController{}

	req, err := http.NewRequest("GET", "/?cep="+input, nil)
	assert.NoError(t, err)

	// Criar um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()

	// Chamar o controlador
	handler := http.HandlerFunc(ctrl.FindData)
	handler.ServeHTTP(rr, req)

	// Verificar o código de status e a resposta
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestInvalidZipCode422(t *testing.T) {
	input := "8"

	//Controller
	ctrl := &controller.FindDataController{}

	req, err := http.NewRequest("GET", "/?cep="+input, nil)
	assert.NoError(t, err)

	// Criar um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()

	// Chamar o controlador
	handler := http.HandlerFunc(ctrl.FindData)
	handler.ServeHTTP(rr, req)

	// Verificar o código de status e a resposta
	assert.Equal(t, 422, rr.Code)
}

func TestValidZipCode(t *testing.T) {
	//mock
	mockUseCase := new(MockUseCase)

	input := "89037501"
	expectedOutput := &dto.ResultOutpurDto{
		Temp_C: 22.7,
		Temp_F: 72.6,
		Temp_K: 295.85,
	}

	//Configura o mock
	mockUseCase.On("FindData", input).Return(expectedOutput, nil)

	//Teste
	result, err := mockUseCase.FindData(nil, input)

	assert.NoError(t, err)
	assert.Equal(t, 22.7, result.Temp_C)
	assert.Equal(t, 72.6, result.Temp_F)
	assert.Equal(t, 295.85, result.Temp_K)
}

func TestNotFoundZipCode(t *testing.T) {
	//mock
	mockUseCase := new(MockUseCase)

	input := "8903750"
	expectedOutput := &dto.ResultOutpurDto{
		Temp_C: 22.7,
		Temp_F: 72.6,
		Temp_K: 295.85,
	}

	expectedErrorOutput := errors.New("inválid zip code")

	//Configura o mock
	mockUseCase.On("FindData", input).Return(expectedOutput, expectedErrorOutput)

	//Teste
	_, err := mockUseCase.FindData(nil, input)

	assert.Error(t, err)
}
