# Changelog

## [2.9.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.8.0...v2.9.0) (2024-03-01)


### Features

* tuningState now respects "mutable" in service.yaml options ([404e794](https://github.com/VU-ASE/pkg-ServiceRunner/commit/404e7942b99dea0c90c3f9408d07936d0eef9ed7))

## [2.8.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.7.2...v2.8.0) (2024-02-28)


### Features

* added package functionality to communicate with the SystemManager ([3847d4e](https://github.com/VU-ASE/pkg-ServiceRunner/commit/3847d4e3e34fed3208175e86e13f3640fb116cb1))

## [2.7.2](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.7.1...v2.7.2) (2024-02-28)


### Bug Fixes

* limit filepath log to last three itesm ([f16010e](https://github.com/VU-ASE/pkg-ServiceRunner/commit/f16010eb7b52a65595877c3712fc3f90ffd2dbc8))

## [2.7.1](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.7.0...v2.7.1) (2024-02-28)


### Bug Fixes

* removed color package ([be0a4e2](https://github.com/VU-ASE/pkg-ServiceRunner/commit/be0a4e2c24a37561ab375a7435fe010047d263c3))

## [2.7.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.6.0...v2.7.0) (2024-02-28)


### Features

* added service prefix to stdout log messages ([39d3b52](https://github.com/VU-ASE/pkg-ServiceRunner/commit/39d3b52378d7cf9feca734133ac64394e6e2fcb3))

## [2.6.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.5.0...v2.6.0) (2024-02-28)


### Features

* added caller info to terminal output ([73c7a13](https://github.com/VU-ASE/pkg-ServiceRunner/commit/73c7a131b4ca68ba8952cf9c0bad2cc19c910445))

## [2.5.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.4.1...v2.5.0) (2024-02-28)


### Features

* add caller information to logs ([70366d2](https://github.com/VU-ASE/pkg-ServiceRunner/commit/70366d283bc745aebd8ea01f14bd4d8562728662))
* added support for error messages from System Manager ([66e1202](https://github.com/VU-ASE/pkg-ServiceRunner/commit/66e12025e95db8ffe865fb1fad09784bb28de1fb))

## [2.4.1](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.4.0...v2.4.1) (2024-02-28)


### Bug Fixes

* not relying on env var anymore, default is tcp://localhost:1337 ([c230e64](https://github.com/VU-ASE/pkg-ServiceRunner/commit/c230e645cbbc82ec1d731a27be4596043fcc3751))

## [2.4.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.3.1...v2.4.0) (2024-02-08)


### Features

* share dependencies with SystemManager ([6f7b02e](https://github.com/VU-ASE/pkg-ServiceRunner/commit/6f7b02e0040090395d0456a6f12db6ada5889cd2))

## [2.3.1](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.3.0...v2.3.1) (2024-02-08)


### Bug Fixes

* correct pkg-CommunicationDefinitions import ([bf132c4](https://github.com/VU-ASE/pkg-ServiceRunner/commit/bf132c4290196798b3d8975351e44a5b0158d738))

## [2.3.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.2.1...v2.3.0) (2024-02-08)


### Features

* implemented dynamic ServiceStatus updates ([62e9532](https://github.com/VU-ASE/pkg-ServiceRunner/commit/62e95322dd7be6a4c7e88e40da2c25816ce09079))

## [2.2.1](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.2.0...v2.2.1) (2024-02-07)


### Bug Fixes

* updated tests ([855b964](https://github.com/VU-ASE/pkg-ServiceRunner/commit/855b964d8db438d0b9b42cb94008ce077d18b8a7))

## [2.2.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.1.0...v2.2.0) (2024-02-07)


### Features

* added support for dynamic and static options ([2b2bf54](https://github.com/VU-ASE/pkg-ServiceRunner/commit/2b2bf549d0ce92a03e59893ad738cd18f24fc7f2))

## [2.1.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.0.1...v2.1.0) (2024-02-06)


### Features

* added support for ServiceList requests ([bc2d58d](https://github.com/VU-ASE/pkg-ServiceRunner/commit/bc2d58d4ddd8adf1cebdeaa5396fc334ed4b8795))

## [2.0.1](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v2.0.0...v2.0.1) (2024-02-06)


### Bug Fixes

* moved all dependencies to v2 ([b4cb3e9](https://github.com/VU-ASE/pkg-ServiceRunner/commit/b4cb3e924e601164ac513b823b494e4f66f1a284))

## [2.0.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v1.4.0...v2.0.0) (2024-02-06)


### ⚠ BREAKING CHANGES

* automatically fetch SystemManager broadcast address

### Features

* automatically fetch SystemManager broadcast address ([6867219](https://github.com/VU-ASE/pkg-ServiceRunner/commit/686721950d1563dfe4545655e2deb3a6ea0dfecb))

## [1.4.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v1.3.0...v1.4.0) (2024-02-06)


### Features

* enabled manual service registration disabling ([2f145e6](https://github.com/VU-ASE/pkg-ServiceRunner/commit/2f145e6ac32eed1e357cd4f53c5a0d09626df632))
* rewrite dependency addresses automatically ([8c1c7a9](https://github.com/VU-ASE/pkg-ServiceRunner/commit/8c1c7a91c39aad9420c8a3dfac89d7e76082c295))

## [1.3.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v1.2.1...v1.3.0) (2024-01-25)


### Features

* replaced git submodules with CommunicationDefinitions package ([bd89edb](https://github.com/VU-ASE/pkg-ServiceRunner/commit/bd89edbc5bc9162f04143dea3d84b0f54c2993fc))

## [1.2.1](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v1.2.0...v1.2.1) (2024-01-25)


### Bug Fixes

* typo in systemmanager protobuf message namespace ([9d8b4bd](https://github.com/VU-ASE/pkg-ServiceRunner/commit/9d8b4bdb6615fc3f80add1fa5117d9252f993c1d))

## [1.2.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v1.1.0...v1.2.0) (2024-01-25)


### Features

* autoreplace localhost with asterisk in own output definitions ([a585f1b](https://github.com/VU-ASE/pkg-ServiceRunner/commit/a585f1bcf415240a53bb0867fa2353a06035e34d))

## [1.1.0](https://github.com/VU-ASE/pkg-ServiceRunner/compare/v1.0.0...v1.1.0) (2024-01-24)


### Features

* moved protobuf messages into own namespace ([05d6868](https://github.com/VU-ASE/pkg-ServiceRunner/commit/05d686889bccf62b1c3a1566d58c16be71751151))

## 1.0.0 (2024-01-24)


### Features

* added CI ([977c2e5](https://github.com/VU-ASE/pkg-ServiceRunner/commit/977c2e530cac8ffe68100a2a5047e2d8858332e4))
* added NON_EXISTENT service status ([4ef5953](https://github.com/VU-ASE/pkg-ServiceRunner/commit/4ef5953c09b134286c60e8f3e3f8f894d4d9e500))
* added service registration and dependency resolving ([1b52864](https://github.com/VU-ASE/pkg-ServiceRunner/commit/1b52864a8e30f4cbc43184981d074ea98fcc14e6))
* added support for tuning states ([f05b831](https://github.com/VU-ASE/pkg-ServiceRunner/commit/f05b831311917a5ecd3668f775eefc3e5c049725))
* added tests ([68a9887](https://github.com/VU-ASE/pkg-ServiceRunner/commit/68a98875827fb8f90cb18a871f399c4e272c44d2))
* basic dependency resolving ([2fcd980](https://github.com/VU-ASE/pkg-ServiceRunner/commit/2fcd980e226f6495fd4ad53cbdf6be09fdfc959d))
* initial code ([caf1e2c](https://github.com/VU-ASE/pkg-ServiceRunner/commit/caf1e2ccfc798d8d7de4e6d90d5c8ec819e59473))
* package README ([f1e2664](https://github.com/VU-ASE/pkg-ServiceRunner/commit/f1e26647c4db6a90243a961d5dca66c83ce1c462))


### Bug Fixes

* avoid system manager registering itself ([b1f9d49](https://github.com/VU-ASE/pkg-ServiceRunner/commit/b1f9d49a14ec60ed74a4bfffa78d045367047684))
* better error messages ([379bd71](https://github.com/VU-ASE/pkg-ServiceRunner/commit/379bd71743912470d7cb638959c342e5d34ecf01))
* correct module path ([0f59c66](https://github.com/VU-ASE/pkg-ServiceRunner/commit/0f59c6657a86d6e21344ac5e047713def1fe6c2a))
* disable tuning state updates for system manager ([b93cd15](https://github.com/VU-ASE/pkg-ServiceRunner/commit/b93cd152f8ea434e0b0e10cb7adba854e7c535bb))
* don't fetch sysman details when service is sysman itself ([e351b08](https://github.com/VU-ASE/pkg-ServiceRunner/commit/e351b08ee80b5c6517030bae8fa99b85904e520e))
* Makefile targets ([1e25ac3](https://github.com/VU-ASE/pkg-ServiceRunner/commit/1e25ac378c253d0bfab093815276188bcc502dc1))
* misisng information in README ([f8e958d](https://github.com/VU-ASE/pkg-ServiceRunner/commit/f8e958d4057770db3fe43cfb2f4f147eea442783))
* reintroduced oneof type for SystemManager message ([e7d9a89](https://github.com/VU-ASE/pkg-ServiceRunner/commit/e7d9a895252852a39540ee42d3da9671c7bf44f2))
* zeromq req/rep socket type ([371c48c](https://github.com/VU-ASE/pkg-ServiceRunner/commit/371c48cf4135d8f9b36557a29e72ee8f0bbae883))
