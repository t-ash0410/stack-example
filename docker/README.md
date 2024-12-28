# docker

Docker関連の設定を管理する。

docker composeについては[こちら](https://docs.docker.jp/compose/toc.html)。

## `docker-compose.local.yaml`

### `dev`

Dev Containersにて利用する、開発環境を構築するためのコンテナ。

Microsoftが提供するコンテナ開発向けのイメージを使用している。

そのままだと起動直後にコンテナが終了してしまうので、起動時に`sleep infinity`を実行する。

### `firestore`

ローカル開発向けのDBとして利用するFirestoreのコンテナ。

UIがアクティブなので、起動後に`localhost:4000`にアクセスするとFirestoreをGUIで操作することができる。

### `firestore-test`

ユニットテスト向けFirestoreのコンテナ。

## `docker-compose.ci.yaml`

### `firestore` (For CI)

GitHub Actions上でユニットテストを実行する際に利用されるFirestoreのコンテナ。

> [!IMPORTANT]
> コンテナ開発環境ではdocker composeによりDNSが構築されているため、`firestore:8080`にアクセスすることでFirestoreを利用することができる。
> しかし、GitHub Actionsではコンテナ内でdocker composeを実行するため、DNSは適用されない。従って、ユニットテストを動かす際は`localhost:8080`に対してアクセスする必要がある。
