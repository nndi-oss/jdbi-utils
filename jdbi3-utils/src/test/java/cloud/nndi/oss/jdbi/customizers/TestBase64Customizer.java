package cloud.nndi.oss.jdbi.customizers;

import cloud.nndi.oss.jdbi.HsqldbDatabaseRule;
import org.jdbi.v3.core.Handle;
import org.jdbi.v3.sqlobject.SqlObjectPlugin;
import org.jdbi.v3.sqlobject.customizer.Bind;
import org.jdbi.v3.sqlobject.statement.SqlUpdate;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class TestBase64Customizer {
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
    public void testShouldEncodeToBase64() {
        hsql.getSharedHandle()
                .createUpdate("INSERT INTO people2(first_name, last_name) VALUES (:first_name, :last_name)")
                .bind("first_name", "John")
                .bind("last_name", "Banda")
                .addCustomizer(new Base64Customizer("first_name"))
                .execute();

        String got = hsql.getSharedHandle()
                .createQuery("SELECT first_name FROM people2 LIMIT 1")
                .mapTo(String.class)
                .first();

        assertEquals("Sm9obg==", got);
    }

    @Test
    public void testShouldEncodeToBase64OnDAO() {
        DAO dao = hsql.onDemand(DAO.class);

        dao.insert( "John", "Banda");

        String got = hsql.getSharedHandle()
                .createQuery("SELECT first_name FROM people2 LIMIT 1")
                .mapTo(String.class)
                .first();

        assertEquals("Sm9obg==", got);
    }

    private static interface DAO {
        @SqlUpdate("INSERT INTO people2(first_name, last_name) VALUES (:first_name, :last_name)")
        @BindBase64({"first_name", "last_name"})
        void insert(@Bind("first_name") String firstName, @Bind("last_name") String lastName);
    }
}
