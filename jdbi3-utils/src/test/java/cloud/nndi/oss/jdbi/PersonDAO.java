package cloud.nndi.oss.jdbi;

import cloud.nndi.oss.jdbi.validation.Valid;
import org.jdbi.v3.core.mapper.RowMapper;
import org.jdbi.v3.core.statement.StatementContext;
import org.jdbi.v3.sqlobject.config.RegisterRowMapper;
import org.jdbi.v3.sqlobject.customizer.Bind;
import org.jdbi.v3.sqlobject.customizer.BindBean;
import org.jdbi.v3.sqlobject.customizer.Timestamped;
import org.jdbi.v3.sqlobject.statement.GetGeneratedKeys;
import org.jdbi.v3.sqlobject.statement.SqlQuery;
import org.jdbi.v3.sqlobject.statement.SqlUpdate;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;

/**
 * SqlObject for tests
 */
@RegisterRowMapper(PersonDAO.PersonRowMapper.class)
public interface PersonDAO {
    @GetGeneratedKeys
    @SqlUpdate("INSERT INTO people(id, firstName, lastName, email, created, modified) VALUES (:p.id, :p.firstName, :p.lastName, :p.email, :now, :now)")
    @Timestamped
    int insert(@BindBean("p") @Valid Person person);

    @SqlUpdate("INSERT INTO people(id, firstName, lastName, email, created, modified) VALUES (:p.id, :p.firstName, :p.lastName, :p.email, :createdAt, :createdAt)")
    @Timestamped("createdAt")
    int insertWithCustomTimestampFields(@BindBean("p") Person person);

    @SqlUpdate("UPDATE people SET firstName = :p.firstName, lastName = :p.lastName, email = :p.email, modified= :now WHERE id = :p.id")
    @Timestamped
    int updatePerson(@BindBean("p") Person person);

    @SqlUpdate("UPDATE people SET email=:p.email WHERE id=:p.id")
    void updateEmail(@BindBean("p") @Valid(groups = Person.EmailUpdate.class) Person person);

    @SqlQuery("SELECT id, firstName, lastName, email, created, modified from people ORDER BY id")
    List<Person> findAll();

    @SqlQuery("SELECT id, firstName, lastName, email, created, modified from people WHERE id=:id")
    Person get(@Bind("id") int id);

    class PersonRowMapper implements RowMapper<Person> {

        @Override
        public Person map(ResultSet resultSet, StatementContext statementContext) throws SQLException {
            Person person = new Person(resultSet.getString("firstName"), resultSet.getString("lastName"), resultSet.getString("email"));
            person.setId(resultSet.getInt("id"));
            person.setCreated(resultSet.getTimestamp("created"));
            person.setModified(resultSet.getTimestamp("modified"));
            return person;
        }
    }
}
