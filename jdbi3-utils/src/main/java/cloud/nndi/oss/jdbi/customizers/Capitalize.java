package cloud.nndi.oss.jdbi.customizers;

import org.jdbi.v3.sqlobject.customizer.SqlStatementCustomizer;
import org.jdbi.v3.sqlobject.customizer.SqlStatementCustomizerFactory;
import org.jdbi.v3.sqlobject.customizer.SqlStatementCustomizingAnnotation;
import org.jdbi.v3.sqlobject.customizer.SqlStatementParameterCustomizer;

import java.lang.annotation.*;
import java.lang.reflect.Method;
import java.lang.reflect.Parameter;
import java.lang.reflect.Type;

@Target({ElementType.METHOD})
@Retention(RetentionPolicy.RUNTIME)
@SqlStatementCustomizingAnnotation(Capitalize.Factory.class)
public @interface Capitalize {
    String[] value();

    class Factory implements SqlStatementCustomizerFactory {
        @Override
        public SqlStatementCustomizer createForMethod(Annotation annotation, Class<?> sqlObjectType, Method method) {
            return customizer((Capitalize) annotation);
        }

        @Override
        public SqlStatementParameterCustomizer createForParameter(Annotation annotation, Class<?> sqlObjectType, Method method, Parameter param, int index, Type parameterType) {
            return (stmt, arg) -> customizer((Capitalize) annotation);
        }

        private SqlStatementCustomizer customizer(Capitalize config) {
            return (stmt) -> stmt.addCustomizer(new CapitalizeCustomizer(config.value()));
        }
    }
}
