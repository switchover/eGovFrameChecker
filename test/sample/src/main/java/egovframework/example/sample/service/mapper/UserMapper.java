package egovframework.example.sample.service.mapper;

import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.factory.Mappers;

@Mapper(componentModel = "spring") // 스프링 빈으로 사용 시
public interface UserMapper {

    // 스프링을 사용하지 않는 경우 인스턴스 접근
    UserMapper INSTANCE = Mappers.getMapper(UserMapper.class);

    @Mapping(source = "emailAddress", target = "email") // 필드명이 다를 때 매핑
    UserDto toDto(User user);

    @Mapping(source = "email", target = "emailAddress")
    User toEntity(UserDto userDto);
}
