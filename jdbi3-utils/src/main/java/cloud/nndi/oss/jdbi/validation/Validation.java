package cloud.nndi.oss.jdbi.validation;

import org.hibernate.validator.HibernateValidator;

import javax.validation.ConstraintViolation;
import javax.validation.ValidationException;
import javax.validation.Validator;
import javax.validation.ValidatorFactory;
import java.util.Collections;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Collectors;

public final class Validation {
    static final ValidatorFactory validatorFactory = javax.validation.Validation
                    .byProvider( HibernateValidator.class )
                    .configure()
                    .buildValidatorFactory();

    /**
     * Validates arguments via Hibernate validator and throws an {@link ValidationException}
     * if the object fails validation otherwise returns the object itself
     */
    public static Object valid(Object object) {
        Validation.throwOnFailedValidation(object);
        return object;
    }

    /**
     * Validates a value and throws exception if there were any validation errors.
     *
     * @param object object to validatee
     * @param groups validation groups
     * @param <T> type parameter
     * @throws ValidationException thrown when validation result has errors
     */
    public static <T> void throwOnFailedValidation(T object, Class<?>... groups) throws ValidationException {
        StringBuilder helper = new StringBuilder();
        Map<String, String> validationErrors = validate(object, groups);
        if (! validationErrors.isEmpty()) {
            validationErrors.forEach((key, value) -> helper.append(key)
                .append("=")
                .append(value)
                .append(" "));
            throw new ValidationException("Entity contains validation errors. Errors: " + helper);
        }
    }


    /**
     * Validate an object of a type and return a map of the errors.
     * The keys are the names of the properties with validation errors.
     *
     * @param value object to validate
     * @param groups validation groups
     * @param <T> type param
     * @return validation errors
     */
    public static <T> Map<String, String> validate(T value, Class<?>... groups) {
        Validator validator = validatorFactory.getValidator();
        if (Objects.nonNull(groups) && groups.length > 0) {
            return validator.validate(value, groups).stream()
                    .collect(Collectors.toMap(cv -> cv.getPropertyPath().toString(), ConstraintViolation::getMessage));
        }
        return validator.validate(value).stream()
                .collect(Collectors.toMap(cv -> cv.getPropertyPath().toString(), ConstraintViolation::getMessage));
    }
}
