package cloud.nndi.oss.jdbi.customizers;

import cloud.nndi.oss.jdbi.HsqldbDatabaseRule;
import cloud.nndi.oss.jdbi.customizers.CapitalizeCustomizer;
import org.jdbi.v3.core.Handle;
import org.jdbi.v3.sqlobject.SqlObjectPlugin;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static cloud.nndi.oss.jdbi.customizers.CapitalizeCustomizer.capitalize;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class TestCapitalizeCustomizer {

    public HsqldbDatabaseRule hsql = new HsqldbDatabaseRule();


    @BeforeEach
    public void beforeEach() throws Exception{
        hsql.before();
        hsql.getJdbi().installPlugin(new SqlObjectPlugin());
        Handle h = hsql.getSharedHandle();
        h.execute("create table people2(id identity primary key, first_name varchar(50), last_name varchar(50));");
    }

    @AfterEach
    public void  afterEach() {
        hsql.after();
    }

    @Test
    public void testShouldIncrementColumnToOne() {
        hsql.getSharedHandle()
            .createUpdate("INSERT INTO people2(first_name, last_name) VALUES (:first_name, :last_name)")
            .bind("first_name", "John William")
            .bind("last_name", "Banda")
            .addCustomizer(capitalize("first_name", "last_name"))
            .execute();

        String got = hsql.getSharedHandle()
            .createQuery("SELECT first_name + ' ' + last_name FROM people2 LIMIT 1")
            .mapTo(String.class)
            .first();

        assertEquals("JOHN WILLIAM BANDA", got);
    }
}
