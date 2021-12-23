package cloud.nndi.oss.jdbi.locator;


import org.junit.jupiter.api.Test;

import java.io.StringReader;
import java.util.Map;
import java.util.Scanner;

import static org.junit.jupiter.api.Assertions.*;

class TestSqlQueriesFileParser {

    @Test
    public void testCanParseSqlQueries() throws Exception {

        String sqlQueries = """
        -- name: exampleQuery
        SELECT * FROM people
        
        -- name: anotherQuery
        SELECT * FROM people
        
        -- name: an invalid query name
        SELECT * FROM people
        
        -- name: anEmptyQuery
        
        -- name: aQuerySplitOverLines
        SELECT * FROM people
        WHERE firstName := firstName AND
        lastName := lastName
        
        ORDER BY id 
        LIMIT 1
        """;

        SqlQueriesFileParser reader = new SqlQueriesFileParser();
        StringReader stringReader = new StringReader(sqlQueries);

        Map<String, String> queries = reader.parseQueries(new Scanner(stringReader));

        queries.size();

        assertTrue(queries.containsKey("exampleQuery"));
        assertTrue(queries.containsKey("anotherQuery"));
        assertTrue(queries.containsKey("anEmptyQuery"));
        assertTrue(queries.containsKey("aQuerySplitOverLines"));
        assertFalse(queries.containsKey("an invalid query name"));

        assertEquals("SELECT * FROM people\n", queries.get("exampleQuery"));
        assertEquals("SELECT * FROM people\n", queries.get("anotherQuery"));
        assertEquals("""
        SELECT * FROM people
        WHERE firstName := firstName AND
        lastName := lastName
        ORDER BY id 
        LIMIT 1
        """, queries.get("aQuerySplitOverLines"));
    }

}