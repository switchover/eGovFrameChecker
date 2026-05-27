package egovframework.example.sample.service.impl;

import java.util.List;

import org.springframework.data.jpa.repository.JpaRepository;

//@Repository
public interface SampleRepository extends JpaRepository<SampleEntity, Long> {
    // 쿼리 메서드 자동 생성 예시
    List<SampleEntity> findByName(String name);
}
