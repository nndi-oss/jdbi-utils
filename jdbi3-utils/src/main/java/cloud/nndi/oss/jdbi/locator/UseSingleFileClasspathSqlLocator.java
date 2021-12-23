package cloud.nndi.oss.jdbi.locator;

public @interface UseSingleFileClasspathSqlLocator {
    /**
     * Name of the SQL file, defaults to the SqlObject's name
     * @return name of the sql file
     */
    String name() default "";
}
