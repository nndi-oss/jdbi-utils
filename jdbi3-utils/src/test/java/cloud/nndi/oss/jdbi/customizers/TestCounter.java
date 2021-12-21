package cloud.nndi.oss.jdbi.customizers;

import cloud.nndi.oss.jdbi.HsqldbDatabaseRule;
import cloud.nndi.oss.jdbi.customizers.Counter;
import cloud.nndi.oss.jdbi.customizers.CounterCustomizer;
import org.jdbi.v3.core.Handle;
import org.jdbi.v3.sqlobject.SqlObjectPlugin;
import org.jdbi.v3.sqlobject.customizer.BindBean;
import org.jdbi.v3.sqlobject.statement.SqlUpdate;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

/**
 * Tests for the CounterCustomizer customizer
 */
public class TestCounter {

    public HsqldbDatabaseRule hsql = new HsqldbDatabaseRule();

    public PostDAO dao;

    @BeforeEach
    public void beforeEach() throws Exception{
        hsql.before();
        hsql.getJdbi().installPlugin(new SqlObjectPlugin());
        Handle h = hsql.getSharedHandle();
        h.execute("create table posts(id identity primary key, content varchar(140), user_id integer);");
        h.execute("create table users(id identity primary key, posts_count integer);");
        h.execute("INSERT INTO users(id, posts_count) VALUES (1, 0);");
        dao = hsql.onDemand(PostDAO.class);
    }

    @AfterEach
    public void  afterEach() {
        hsql.after();
    }

    @Test
    public void testShouldIncrementColumnToOne() {
        hsql.getSharedHandle()
            .createUpdate("INSERT INTO posts(content, user_id) VALUES (:content, :author_id)")
            .bind("content", "Yay! Post content!")
            .bind("author_id", 1)
            .addCustomizer(new CounterCustomizer("users", "posts_count", "author_id", "id"))
            .execute();

        int postCount = hsql.getSharedHandle()
                .createQuery("SELECT posts_count FROM users WHERE id = 1")
                .mapTo(Integer.class)
                .first();

        assertEquals(1, postCount);
    }

    @Test
    public void testShouldWorkOnSqlObject() {
        int beforeInsertingPost = hsql.getSharedHandle()
                .createQuery("SELECT posts_count FROM users WHERE id = 1")
                .mapTo(Integer.class)
                .first();

        assertEquals(0, beforeInsertingPost);

        dao.insert(new Post("Woo! Post content!", 1));

        int afterInsertingPost = hsql.getSharedHandle()
                .createQuery("SELECT posts_count FROM users WHERE id = 1")
                .mapTo(Integer.class)
                .first();

        assertEquals(1, afterInsertingPost);
    }

    @Test
    public void testShouldDecrementCounter() {
        dao.insert(new Post("Woo! Post content!", 1));
        dao.insert(new Post("Woo! More post content!", 1));

        dao.delete(new Post(1, 1));

        int afterDeletingPost = hsql.getSharedHandle()
                .createQuery("SELECT posts_count FROM users WHERE id = 1")
                .mapTo(Integer.class)
                .first();

        assertEquals(1, afterDeletingPost);
    }

    public static class Post {
        private long id;
        private String content;
        private long userId;

        public Post(String content, long userId) {
            this.content = content;
            this.userId = userId;
        }

        public Post(long id, long userId) {
            this.id = id;
            this.userId = userId;
        }

        public long getId() {
            return id;
        }

        public String getContent() {
            return content;
        }

        public long getUserId() {
            return userId;
        }
    }

    public interface PostDAO {
        @SqlUpdate("INSERT INTO posts(content, user_id) VALUES (:p.content, :p.userId)")
        @Counter(table = "users", column = "posts_count", binding = "p.userId")
        void insert(@BindBean("p") Post post);

        @SqlUpdate("DELETE FROM posts WHERE id = :p.id")
        @Counter(table = "users", column = "posts_count", binding = "p.userId", decrementing = true)
        void delete(@BindBean("p") Post post);
    }
}
