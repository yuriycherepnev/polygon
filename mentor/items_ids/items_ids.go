// Есть список itemIDs.
// Для каждого идентификатора товара (itemID) необходимо вызвать медленную функцию fetchStock(id) и получить результат.
// При этом одновременно должно выполняться не более workersCount запросов.
// Для каждого полученного результата:
// если возникла ошибка — вызвать функцию processError;
// если результат получен успешно — вызвать функцию saveResult

// *Отмена выполнения функции

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Result struct {
	itemID string
	stock  int
	err    error
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
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	itemsCh := make(chan string)
	resultsCh := make(chan Result)
	wg := &sync.WaitGroup{}

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go s.worker(ctx, itemsCh, resultsCh, wg)
	}

	go func() {
		for _, itemID := range itemIDs {
			select {
			case <-ctx.Done():
				return
			case itemsCh <- itemID:
			}
		}

		close(itemsCh)
	}()

	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	for result := range resultsCh {
		fmt.Println(result)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case result, ok := <-resultsCh:
			if !ok {
				return nil
			}

			if result.err != nil {
				err := s.errorProcessor.ProcessError(ctx, result.itemID, result.err)
				if err != nil {
					cancel()
					return err
				}

				continue
			}

			err := s.resultSaver.SaveResult(ctx, result.itemID, result.stock)
			if err != nil {
				cancel()
				return err
			}
		}
	}
}

func (s *StockService) worker(ctx context.Context, itemsCh chan string, resultsCh chan Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case itemID, ok := <-itemsCh:
			if !ok {
				return
			}
			stock, err := s.fetcher.FetchStock(ctx, itemID)

			result := Result{
				itemID: itemID,
				stock:  stock,
				err:    err,
			}

			select {
			case <-ctx.Done():
				return
			case resultsCh <- result:
			}
		}
	}
}

type fetcher struct{}

func (f *fetcher) FetchStock(ctx context.Context, itemID string) (int, error) {
	time.Sleep(5 * time.Second)
	return 10, nil
}

type errorProcessor struct{}

func (p *errorProcessor) ProcessError(ctx context.Context, itemID string, err error) error {
	fmt.Println("error:", itemID, err)
	return nil
}

type resultSaver struct{}

func (s *resultSaver) SaveResult(ctx context.Context, itemID string, stock int) error {
	fmt.Println("saved:", itemID, stock)
	return nil
}

func main() {
	service := NewStockService(
		&fetcher{},
		&errorProcessor{},
		&resultSaver{},
	)

	err := service.ProcessStocks(context.Background(), []string{"1", "2", "3"}, 2)

	fmt.Println(err)

}
