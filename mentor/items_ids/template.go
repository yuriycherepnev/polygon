// Есть список itemIDs.
// Для каждого идентификатора товара (itemID) необходимо вызвать медленную функцию fetchStock(id) и получить результат.
// При этом одновременно должно выполняться не более workersCount запросов.
// Для каждого полученного результата:
// если возникла ошибка — вызвать функцию processError;
// если результат получен успешно — вызвать функцию saveResult

// *Отмена выполнения функции

package main

/*
import "context"

type Result struct {
	val int
	err error
}

type StockFetcher interface {
	FetchStock(ctx context.Context, itemID string) (int, error)
}

type ErrorProcessor interface {
	ProcessError(ctx context.Context, itemID string, err error) error
}

type ResultSaver interface {
	SaveResult(ctx context.Context, itemID string, stock int) error
}

type StockService struct {
	fetcher        StockFetcher
	errorProcessor ErrorProcessor
	resultSaver    ResultSaver
}

func NewStockService(
	fetcher StockFetcher,
	errorProcessor ErrorProcessor,
	resultSaver ResultSaver,
) *StockService {
	return &StockService{
		fetcher:        fetcher,
		errorProcessor: errorProcessor,
		resultSaver:    resultSaver,
	}
}

func (s *StockService) ProcessStocks(
	ctx context.Context,
	itemIDs []string,
	workersCount int,
) error {
	// implementation
	return nil
}
*/
