package cloud.nndi.oss.jdbi;

import org.jdbi.v3.sqlobject.locator.UseClasspathSqlLocator;
import org.jdbi.v3.sqlobject.statement.SqlQuery;

@UseClasspathSqlLocator
public interface ExampleDAO {
    @SqlQuery
    int selectOne();
}
