/**
 * @name functionsList
 * @kind Show functions in a file
 * @problem.severity Information
 * @id @0xCVEer
 */

import go

from Function fun, File file
where file.getBaseName() = "main.go"
select fun.getDeclaration()
