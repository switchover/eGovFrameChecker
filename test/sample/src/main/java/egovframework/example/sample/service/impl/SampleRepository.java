package egovframework.example.sample.service.impl;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

//@Repository
public interface SampleRepository extends JpaRepository<MemberEntity, Long> {
    // 쿼리 메서드 자동 생성 예시
    List<Member> findByName(String name);
}
