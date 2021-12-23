package cloud.nndi.oss.jdbi.locator;

import java.io.InputStream;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

class SqlQueriesFileParser {
    private static final Pattern reTag = Pattern.compile("^\s*--\s*(?<tagName>.+)\s*:\s*(?<methodName>.+)");
    private static final Pattern reComment = Pattern.compile("^\\s*--\\s*(.*)");
    private static final Pattern reMethodName = Pattern.compile("\s?\\w+\s?");

    private boolean skipInvalidTags = true;
    private boolean removingComments = false;

    public SqlQueriesFileParser() {
        this(true, false);
    }

    public SqlQueriesFileParser(boolean skipInvalidTags, boolean removingComments) {
        this.skipInvalidTags = skipInvalidTags;
        this.removingComments = removingComments;
    }

    public Map<String, String> parseQueries(InputStream inputStream) throws Exception {
        Map<String, String> queriesMap = new HashMap<>();
        return parseQueries(new Scanner(inputStream));
    }

    public Map<String, String> parseQueries(Scanner scanner) throws Exception {
        Map<String, String> queriesMap = new HashMap<>();
        try (Scanner sc = scanner) {
            StringBuilder sb = null; // for the query
            String currentLine = null;
            String currentTag = null;
            /** This loop does the following things
             * 0. Reads each line in the input
             * 1. Normalize messages so that each appears as one line
             * 2. Skips <media omitted> lines in the input
             */
            while(sc.hasNextLine()) {
                currentLine = sc.nextLine();
                if (currentLine.trim().isBlank() || currentLine.isEmpty()) {
                    continue; // Skip empty lines
                }

                Matcher tagMatcher = reTag.matcher(currentLine);
                if (tagMatcher.matches()) {
                    if (currentTag != null) {
                        queriesMap.put(currentTag, sb.toString());
                        currentTag = null;
                    }
                    if (! "name".equalsIgnoreCase(tagMatcher.group("tagName"))) {
                        continue;
                    }
                    currentTag = tagMatcher.group("methodName").trim();
                    if (skipInvalidTags && !reMethodName.matcher(currentTag).matches()) {
                        currentTag = null;
                        continue;
                    }
                    queriesMap.put(currentTag, "");
                    sb = new StringBuilder(); // reset the string builder
                    continue;
                }

                if (sb == null) {
                    continue;
                }
                if (removingComments && reComment.matcher(currentLine).find()) {
                    continue;
                }
                // we append comments and queries verbatim and append a line after each
                sb.append(currentLine).append("\n");
                if (! sc.hasNextLine()) {
                    queriesMap.put(currentTag, sb.toString());
                }
            }

            return queriesMap;
        }
    }
}
