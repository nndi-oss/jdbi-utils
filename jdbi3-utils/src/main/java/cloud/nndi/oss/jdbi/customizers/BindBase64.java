package cloud.nndi.oss.jdbi.customizers;

import org.jdbi.v3.sqlobject.customizer.SqlStatementCustomizer;
import org.jdbi.v3.sqlobject.customizer.SqlStatementCustomizerFactory;
import org.jdbi.v3.sqlobject.customizer.SqlStatementCustomizingAnnotation;
import org.jdbi.v3.sqlobject.customizer.SqlStatementParameterCustomizer;

import java.lang.annotation.*;
import java.lang.reflect.Method;
import java.lang.reflect.Parameter;
import java.lang.reflect.Type;

@Target({ElementType.TYPE, ElementType.METHOD})
@Retention(RetentionPolicy.RUNTIME)
@SqlStatementCustomizingAnnotation(BindBase64.Factory.class)
public @interface BindBase64 {
    String[] value() ;

    class Factory implements SqlStatementCustomizerFactory {
        @Override
        public SqlStatementCustomizer createForMethod(Annotation annotation, Class<?> sqlObjectType, Method method) {
            return customizer((BindBase64) annotation);
        }

        @Override
        public SqlStatementParameterCustomizer createForParameter(Annotation annotation, Class<?> sqlObjectType, Method method, Parameter param, int index, Type parameterType) {
            return (stmt, arg) -> customizer((BindBase64) annotation);
        }

        private SqlStatementCustomizer customizer(BindBase64 config) {
            return (stmt) -> stmt.addCustomizer(new Base64Customizer(config.value()));
        }
    }
}
