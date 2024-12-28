# .devcontainer

Dev Containers関連の設定を管理している。

Dev Containersについては[公式ドキュメント](https://code.visualstudio.com/docs/devcontainers/containers)を参照してほしいが、VSCodeを使ってContainer上での開発を可能にするツールである。

workspaceのトップレベルに`.devcontainer/devcontainer.json`がある場合、Dev Containersはそれをデフォルトの設定ファイルとして認識してくれる。

## `devcontainer.json`

Dev Containersの設定値を記述するJSONファイル。

各項目に対して適切に値をセットすることで、Dev Containersがそれを認識して開発環境用のコンテナをビルドし、VSCodeから参照可能な状態にしてくれる。

各項目の詳細は[こちら](https://containers.dev/implementors/json_reference/)に記述されているが、以下はこのプロジェクトにおいて特に重要な項目の紹介である。

### `dockerComposeFile`

開発用コンテナとして扱うイメージが記述されたdocker-composeファイルのパスを記述する。

Dev Containersは以下の方法でコンテナを設定可能である。

- image名を直接指定
- DockerFileを指定
- docker-composeファイルを指定

今回は開発用途以外にも複数のコンテナを起動する必要があったため、docker-composeファイルを指定している。

### `service`

`dockerComposeFile`が指定されている場合に必須。docker-composeファイル内のどのサービスを開発用コンテナとして扱うのかを指定する。

docker-composeファイル内に実在するサービス名を指定しなければならない。

### `customizations.vscode.settings`

Dev Containersを利用したときだけ有効となるVSCodeの設定値を記述できる。

開発者間での設定差異を無くすことができる。

### `customizations.vscode.extensions`

Dev Containersを利用したときだけインストールされるVSCodeの拡張機能を記述できる。

開発者間で共通した拡張機能を利用することが可能になる。

### `containerEnv`

コンテナの環境変数を指定できる。

`${localEnv:XXX}`など指定することでローカルPCの環境変数を参照することも可能。

詳細は[こちら](https://code.visualstudio.com/remote/advancedcontainers/environment-variables)。

### `features`

コンテナ起動時にインストールするツールを指定することができる。

コンテナ内に含まれていないものの開発上必要となるツールをここに記述することで、コンテナをビルドするタイミングでそのインストールをDev Containersが代行してくれる。

利用可能なツールの一覧は[こちら](https://containers.dev/features)。

### `onCreateCommand`, `postStartCommand`

Dev Containersでは、コンテナを起動するライフサイクル上にスクリプトを注入することができる。詳細は[こちら](https://containers.dev/implementors/json_reference/#lifecycle-scripts)。

onCreateCommandとpostStartCommandはその一種で、それぞれ以下役割である。

- onCreateCommand: コンテナが初めて開始された直後にそのコンテナ内で実行されるスクリプト
  - `on-create.sh`を実行する。
- postStartCommand: コンテナが正常に起動するたびに実行されるスクリプト
  - コンテナ上でgitコマンドを利用するために、`git config --global --add safe.directory`を実行する。

## on-create.sh

Dev Containersを使って開発環境をセットアップした段階で、ストレスなく開発に着手できるように以下を行っている。

- 各パッケージにおける依存関係のインストール
