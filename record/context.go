// Stuff related to recorder options settings, what to record and how.

package record

//import (
//	"github.com/lesovsky/pgcenter/lib/stat"
//)
//
//// Setup method performs context's setup - select queries and adjusts then depending on Postgres version
//func (o *RecordOptions) Setup(pginfo stat.PgInfo) {
//	o.contextList = stat.ContextList{
//		stat.DatabaseView:            &stat.PgStatDatabaseUnit,
//		stat.ReplicationView:         &stat.PgStatReplicationUnit,
//		stat.TablesView:              &stat.PgStatTablesUnit,
//		stat.IndexesView:             &stat.PgStatIndexesUnit,
//		stat.SizesView:               &stat.PgTablesSizesUnit,
//		stat.FunctionsView:           &stat.PgStatFunctionsUnit,
//		stat.ProgressVacuumView:      &stat.PgStatProgressVacuumUnit,
//		stat.ProgressClusterView:     &stat.PgStatProgressClusterUnit,
//		stat.ProgressCreateIndexView: &stat.PgStatProgressCreateIndexUnit,
//		stat.ActivityView:            &stat.PgStatActivityUnit,
//		stat.StatementsTimingView:    &stat.PgSSTimingUnit,
//		stat.StatementsGeneralView:   &stat.PgSSGeneralUnit,
//		stat.StatementsIOView:        &stat.PgSSIoUnit,
//		stat.StatementsTempView:      &stat.PgSSTempUnit,
//		stat.StatementsLocalView:     &stat.PgSSLocalUnit,
//	}
//
//	// Adjust queries depending on Postgres version
//	// Pass truncation limit to query adjustment options
//	o.sharedOptions.PgSSQueryLen = o.TruncLimit
//
//	//
//	o.contextList.AdjustQueries(pginfo)
//	o.sharedOptions.Adjust(pginfo, "record")
//}
