# eGovFrameChecker
eGovFrame Compatibility Checker

[![Build](https://github.com/switchover/eGovFrameChecker/actions/workflows/build.yml/badge.svg)](https://github.com/switchover/eGovFrameChecker/actions/workflows/build.yml)

## Overview
eGovFrameChecker는 전자정부 표준프레임워크(eGovFrame) 호환성 가이드 기준에 맞는지를 점검하는 CLI(Command Line Interface) 도구입니다.

<img width="800" src="./doc/checklist.gif" alt="checklist 명령 예시"/>

### Features
eGovFrameChecker는 관련 정보(표준프레임워크 사용 버전, 각 컴포넌트 사용 규칙 등)를 입력해 호환성 가이드 기준 여부를 판단하는 `checklist` 기능과 
전체 소스를 대상으로 각 레이어(`Controller`, `Service`, `Repository`) 기준을 확인하는 `inspect` 기능을 제공합니다.

세부적인 활용 방법은 다음과 같이 `-h` 또는 `--help` 도움말 옵션을 통해 확인할 수 있습니다.
```shell
$ ./egovchecker --help
eGovFrameChecker is a cli tool for checking eGovFrame compatibility. It checks whether the architecture criteria are met.

Usage:
  egovchecker [command]

Available Commands:
  checklist     Check the checklist for eGovFrame compatibility verification
  defaultconfig Write default 'config.ini' file
  help          Help about any command
  inspect       Inspect eGovFrame compatibility architecture criteria
  version       Print the version of eGovFrameChecker

Flags:
      --config string   config file (default is $HOME/config.ini. Or ./config.ini is used)
  -h, --help            help for egovchecker

Use "egovchecker [command] --help" for more information about a command.
```
※ Windows 환경에서는 `egovchecker.exe`를 사용하면 됩니다.
```shell
> egovchecker.exe -h
```

### Docker 실행
Docker를 통해서도 다음과 같이 실행할 수 있습니다.
```shell
$ docker run --rm switchover/egovchecker:latest -h
```
Docker를 통한 다른 명령 실행은 [Docker Hub](https://hub.docker.com/r/switchover/egovchecker) 페이지를 참고하세요.

## `checklist` Command
`checklist` 명령은 표준프레임워크 호환성 가이드에 맞는지를 제시된 체크리스트를 통해 확인할 수 있습니다.
```shell
$ ./egovchecker checklist
Use the arrow keys to navigate: ↓ ↑ → ← 
? eGovFrame version used: 
    v4.3
  ▸ v4.2
    v4.1
    v4.0
↓   v3.10
```
제시된 체크리스트에 대해 화살표를 통해 선택하거나, 직접 입력하면 다음과 같이 최종 결과를 확인할 수 있습니다.
```
------------------------------------------------------------
 ______    _______  _______  __   __  ___      _______
|    _ |  |       ||       ||  | |  ||   |    |       |
|   | ||  |    ___||  _____||  | |  ||   |    |_     _|
|   |_||_ |   |___ | |_____ |  |_|  ||   |      |   |
|    __  ||    ___||_____  ||       ||   |___   |   |
|   |  | ||   |___  _____| ||       ||       |  |   |
|___|  |_||_______||_______||_______||_______|  |___|  
------------------------------------------------------------
eGovFrame Compatibility: Compatible (level: 2)
  - Version criteria: Satisfied
  - Configuration criteria: Satisfied
  - MVC criteria: Satisfied
  - Service criteria: Satisfied
  - Data access criteria: Satisfied
------------------------------------------------------------
```

참고로 표준프레임워크 버전에 따른 기준은 [criteria.json](./internal/criteria/assets/criteria.json)에서 확인할 수 있습니다.


## `inspect` Command
`inspect` 명령은 전체 소스를 대상으로 각 레이어(`Controller`, `Service`, `Repository`) 기준을 확인할 수 있습니다.

우선 다음과 같이 `inspect` 명령에 대한 세부 옵션을 확인합니다.
```shell
$ ./egovchecker inspect --help 
Inspect eGovFrame compatibility architecture criteria and save related data.

Usage:
  egovchecker inspect [flags]

Flags:
  -h, --help              help for inspect
  -o, --output            Output result to CSV file
  -p, --packages string   Packages to inspect with comma separated
  -s, --skip              Skip file error
  -t, --target string     Target directory to inspect
  -v, --verbose           Verbose output

Global Flags:
      --config string   config file (default is $HOME/config.ini. Or ./config.ini is used)
```
세부 옵션은 다음과 같습니다.

| 옵션                 | 설명                       | 비고 |
|--------------------|--------------------------|----|
| `-t`, `--target`   | 검사할 대상 디렉토리를 지정합니다.      | 필수 |
| `-p`, `--packages` | 검사할 패키지를 콤마로 구분하여 입력합니다. | 필수 |
| `-o`, `--output`   | 결과를 CSV 파일들로 출력합니다.      |    |
| `-s`, `--skip`     | 파일 처리 오류를 무시하고 진행합니다.    |    |
| `-v`, `--verbose`  | 자세한 정보를 출력합니다.           |    |

참고로, `-o` 옵션을 지정하면 `controllers.csv`, `services.csv`, `repositories.csv` 3개의 csv(comma-separated values) 파일이 현 디렉토리에 저장되며,
각 파일별 기준 충족 여부를 확인할 수 있습니다. (해당 내용은 실행 화면에도 표시됨)

### Example
다음은 샘플 게시판 예제에 대한 점검을 수행 예제입니다.
```shell
$ ./egovchecker inspect -t ./test/sample -p egovframework -o
2025/01/12 17:54:42 Total java files: 11
2025/01/12 17:54:42 --------------------------------------------------------------------------------
2025/01/12 17:54:42 ✔ Controller(1 file) is OK.
2025/01/12 17:54:42 --------------------------------------------------------------------------------
2025/01/12 17:54:42 Output file written: controllers.csv
2025/01/12 17:54:42 --------------------------------------------------------------------------------
2025/01/12 17:54:42 ✔ Service(1 file) is OK.
2025/01/12 17:54:42 --------------------------------------------------------------------------------
2025/01/12 17:54:42 Output file written: services.csv
2025/01/12 17:54:42 --------------------------------------------------------------------------------
2025/01/12 17:54:42 ✔ Repository(1 file) is OK.
2025/01/12 17:54:42 --------------------------------------------------------------------------------
2025/01/12 17:54:42 Output file written: repositories.csv
```


## Customizing `inspect` Command 
`inspect` 명령은 내부적으로 다음과 같은 점검 규칙을 사용하며, 별도로 규칙을 수정해 적용할 수 있습니다.

기본 설정은 [config.ini](./internal/config/assets/config.ini) 파일을 통해 확인 가능하며 다음과 같이 구성됩니다.
```ini
# This is a default configuration file for the program

[controller]
fileNameGlobPattern = *Controller
classAnnotations = @Controller,@RestController
methodAnnotations = @RequestMapping,@GetMapping,@PostMapping,@PutMapping,@DeleteMapping

[service]
fileNameGlobPattern = *ServiceImpl
implementation = true
classAnnotations = @Service
superClasses = EgovAbstractServiceImpl

[repository]
fileNameGlobPattern = *{DAO,Mapper}

[repository.ibatis]
classAnnotations = @Repository
superClasses = EgovAbstractDAO,EgovComAbstractDAO

[repository.mybatis]
classAnnotations = @Repository
superClasses = EgovAbstractMapper,EgovComAbstractDAO

[repository.mapper]
classAnnotations = @Mapper
interface = true

[repository.jpa]
classAnnotations = @Repository
interface = true
superClasses = JpaRepository,CrudRepository,PagingAndSortingRepository

[repository.hibernate]
classAnnotations = @Repository
fieldTypes = HibernateTemplate,EntityManger,EntityManagerFactory,Session,SessionFactory
```
Ini 파일 설정은 3개의 세션(`controller`, `service`, `repository`)과 
`repository` 세션 밑에 5개의 하위 세션(`repository.ibatis`, `repository.mybatis`, `repository.mapper`, `repository.jpa`, `repository.hibernate`)로 구성되어 있습니다.

각 세션은 대상 레이어의 클래스에 대해 다음과 같은 설정들을 선택적으로 가집니다. 
- `fileNameGlobPattern`: 파일이름 Glob 패턴
- `classAnnotations`: 지정되어야 할 클래스 어노테이션 (여러 개는 콤마로 구분되며 or 조건으로 적용)
- `methodAnnotations`: 활용되어야 하는 메소드 어노테이션 (여러 개는 콤마로 구분되며 or 조건으로 적용)
- `implementation`: 인터페이스 구현 여부
- `superClasses`: 상속되어야 하는 클래스 (여러 개는 콤마로 구분되며 or 조건으로 적용)
- `interface`: 인터페이스 여부
- `fieldTypes`: 활용되어야 하는 필드 타입 (여러 개는 콤마로 구분되며 or 조건으로 적용)

### 설정 예시
위 기본 설정에 의하면 Service는 `*ServiceImpl`로 끝나는 파일이며, 
인터페이스를 구현(`implments`)해야 하며, `@Service` 어노테이션 및 `EgovAbstractServiceImpl` 클래스를 상속해야 합니다.

### 기본 규칙 변경 방법
위 설정 규칙을 변경하려면 다음과 같은 명령을 통해 기본 `config.ini` 파일을 생성한 후 수정할 수 있습니다.
```shell
$ ./egovchecker defaultconfig
2025/01/12 18:21:25 Default config file 'config.ini' has been written.
```

혹 현재 디렉토리에 `config.ini` 파일이 존재하면 다음과 같이 출력되고 생성되지 않습니다.
이 경우 별도 옵션(`--overwrite`)을 지정하면 덮어쓸 수 있습니다.
```shell
$ ./egovchecker defaultconfig
2025/01/12 18:22:27 'config.ini' file exists. Use --overwrite flag to overwrite.
$ rm config.ini
$ ./egovchecker defaultconfig
2025/01/12 18:23:00 Default config file 'config.ini' has been written.
```

이제 `config.ini` 파일을 수정한 후 `inspect` 명령을 실행하면 변경된 규칙이 적용됩니다.
(`config.ini` 파일은 프로그램 실행 시 현재 디렉토리 또는 홈 디렉토리에서 찾음)


## Build
### Go Build
다음과 같이 빌드 스크립트를 통해 븰드를 수행할 수 있습니다.
```shell
$ ./build_with_flags.sh -o egovchecker
```
Windows 환경에서는 `build_with_flags.bat` 파일을 사용하며, 
스크립트 내부적으로 Go 언어 버전, 빌드 시간 및 commit 정보를 buildflag로 추가합니다.

### Docker Build
Docker image는 다음과 같이 빌드 및 배포할 수 있습니다.
```shell
$ docker build -t switchover/egovchecker .
$ docker tag switchover/egovchecker switchover/egovchecker:[version]
$ docker login
$ docker push switchover/egovchecker
$ docker push switchover/egovchecker:[version]
```
- tag 부분이 지정되지 않으면 `latest`가 자동 지정됨
- `[version]`은 버전 정보로 `0.1` 등과 같이 버전 정보를 지정합니다.


## Reference
- [표준프레임워크 호환성 가이드](./doc/guidelines_20241114.pdf)
- [ANTLR4 Go 언어 런타임](https://github.com/antlr/antlr4/blob/master/doc/go-target.md)
- [Java ANTLR Grammar file](https://github.com/antlr/grammars-v4/tree/master/java/java)
