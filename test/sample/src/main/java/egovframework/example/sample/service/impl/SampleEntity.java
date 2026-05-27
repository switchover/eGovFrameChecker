package egovframework.example.sample.service.impl;

import jakarta.persistence.*;

@Entity
@Table(name = "samples")
public class SampleEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(nullable = false)
    private String name;

    // Getters, Setters, Constructors
}
