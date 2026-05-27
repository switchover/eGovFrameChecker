package egovframework.example.cmmn;

import java.lang.annotation.*;

@Retention(value = RetentionPolicy.RUNTIME)
@Target(value = ElementType.TYPE)
public @interface CommonMapper {
    String value() default "";
}
