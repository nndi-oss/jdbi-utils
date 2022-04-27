package cloud.nndi.oss.jdbi.customizers;

import org.jdbi.v3.core.argument.Argument;
import org.jdbi.v3.core.statement.StatementContext;
import org.jdbi.v3.core.statement.StatementCustomizer;
import org.jdbi.v3.sqlobject.customizer.SqlStatementCustomizer;
import org.jdbi.v3.sqlobject.customizer.SqlStatementCustomizerFactory;
import org.jdbi.v3.sqlobject.customizer.SqlStatementParameterCustomizer;
import org.slf4j.LoggerFactory;

import java.lang.annotation.Annotation;
import java.lang.reflect.Method;
import java.lang.reflect.Parameter;
import java.lang.reflect.Type;
import java.nio.charset.StandardCharsets;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.Arrays;
import java.util.Base64;
import java.util.Optional;

public class Base64Customizer implements StatementCustomizer {
    private final String[] bindings;
    private final Base64.Encoder base64Encoder = Base64.getEncoder();

    public Base64Customizer(String... bindings) {
        this.bindings = bindings;
    }

    @Override
    public void beforeBinding(PreparedStatement stmt, StatementContext ctx) throws SQLException {
        Arrays.stream(bindings).forEach(binding -> {
            Optional<Argument> bindingVal = ctx.getBinding().findForName(binding, ctx);

            if (bindingVal.isEmpty()) {
                LoggerFactory.getLogger(getClass()).warn("Missing binding '{}'. Cannot Base64 encode", binding);
                return;
            }
            // TODO: Check the type of the binding and read the contents if it's a File, Path or URL?
            ctx.getBinding().addNamed(binding, base64Encoder.encodeToString(String.valueOf(bindingVal.get()).getBytes(StandardCharsets.UTF_8)));
        });
    }
}
