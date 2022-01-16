package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/dajinkuang/dlog"
)

// DBExecContext 数据库操作 ExecContext
func DBExecContext(ctx context.Context, __DB *sql.DB, query string, args ...interface{}) (result sql.Result, err error) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		if err != nil {
			dlog.ErrorContext(ctx, "mysql-DBExecContext-error", "query", query, "args", args, "error", err, "dur/ms", dur)
		} else {
			dlog.InfoContext(ctx, "mysql-DBExecContext-info", "query", query, "args", args, "dur/ms", dur)
		}
	}()
	result, err = __DB.ExecContext(ctx, query, args...)
	return result, err
}

// DBQueryContext 数据库操作 QueryContext
func DBQueryContext(ctx context.Context, __DB *sql.DB, query string, args ...interface{}) (rows *sql.Rows, err error) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		if err != nil {
			dlog.ErrorContext(ctx, "mysql-DBQueryContext-error", "query", query, "args", args, "error", err, "dur/ms", dur)
		} else {
			dlog.InfoContext(ctx, "mysql-DBQueryContext-info", "query", query, "args", args, "dur/ms", dur)
		}
	}()
	rows, err = __DB.QueryContext(ctx, query, args...)
	return rows, err
}

// DBQueryRowContext 数据库操作 QueryRowContext
func DBQueryRowContext(ctx context.Context, __DB *sql.DB, query string, args ...interface{}) (row *sql.Row) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		dlog.InfoContext(ctx, "mysql-DBQueryRowContext-info", "query", query, "args", args, "dur/ms", dur)
	}()
	row = __DB.QueryRowContext(ctx, query, args...)
	return row
}

// TxExecContext 数据库事务操作 ExecContext
func TxExecContext(ctx context.Context, __Tx *sql.Tx, query string, args ...interface{}) (result sql.Result, err error) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		if err != nil {
			dlog.ErrorContext(ctx, "mysql-TxExecContext-error", "query", query, "args", args, "error", err, "dur/ms", dur)
		} else {
			dlog.InfoContext(ctx, "mysql-TxExecContext-info", "query", query, "args", args, "dur/ms", dur)
		}
	}()
	result, err = __Tx.ExecContext(ctx, query, args...)
	return result, err
}

// TxQueryContext 数据库事务操作 QueryContext
func TxQueryContext(ctx context.Context, __Tx *sql.Tx, query string, args ...interface{}) (rows *sql.Rows, err error) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		if err != nil {
			dlog.ErrorContext(ctx, "mysql-TxQueryContext-error", "query", query, "args", args, "error", err, "dur/ms", dur)
		} else {
			dlog.InfoContext(ctx, "mysql-TxQueryContext-info", "query", query, "args", args, "dur/ms", dur)
		}
	}()
	rows, err = __Tx.QueryContext(ctx, query, args...)
	return rows, err
}

// TxQueryRowContext 数据库事务操作 QueryRowContext
func TxQueryRowContext(ctx context.Context, __Tx *sql.Tx, query string, args ...interface{}) (row *sql.Row) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		dlog.InfoContext(ctx, "mysql-TxQueryRowContext-info", "query", query, "args", args, "dur/ms", dur)
	}()
	row = __Tx.QueryRowContext(ctx, query, args...)
	return row
}
