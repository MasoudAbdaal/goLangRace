import go
import semmle.code.go.controlflow.ControlFlowGraph

// Define a class to represent SQL queries
class SQLQuery extends go:CallExpr {
  SQLQuery() {
    exists(go:CallExpr call |
      this = call.getAPrimaryExpr() and
      call.getTarget().getName() = "Query" and
      call.getPackage().getName() = "database/sql"
    )
  }
}

// Define a predicate to find potential SQL injection points
predicate potentialInjectionPoint(SQLQuery query, int index) {
  exists(go:FunctionCallExpr call |
    call.getTarget().getQualifiedName().matches(".*[Ee]xecute.*") and
    call.getAnArgument().getAnArgument(index) and
    query.getArgument(index) = call.getAnArgument().getAnArgument(index) and
    not call.getAnArgument().getAnArgument(index).(go:BasicLit)
  )
}

// Define a predicate to find sink points for potential injections
predicate sink(SQLQuery query, int index) {
  exists(FunctionCallExpr call |
    call.getTarget().getQualifiedName().matches(".*[Ee]xecute.*") and
    call.getAnArgument().getAnArgument(index) = query.getArgument(index) and
    not call.getAnArgument().getAnArgument(index).(go:BasicLit)
  )
}

from
  SQLQuery query,
  int index, string param
where
  potentialInjectionPoint(query, index) and
  sink(query, index) and
  param = query.getArgument(index).getValue().toString() and
  contains(param, ";")
select
  "Potential SQL injection found at index ", index, " in function ", query.getEnclosingFunction().getName(),
  " in file ", query.getEnclosingFunction().getFile().getDeclaredLocation().getFileName()
