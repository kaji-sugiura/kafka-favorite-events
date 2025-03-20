# kafka-favorite-events

## お気に入りイベント非同期集計システム (Kafka × Go × MySQL)

## 概要

このプロジェクトは、Kafka を用いたイベント駆動アーキテクチャをベースに、お気に入りイベントを非同期で処理しランキング集計を行うシステムの PoC です。スケール設計、非同期処理、データ集計、DLQ 設計までを学習目的で実装しています。

## システム構成

- API Server（Go）

- Kafka（Docker Compose 上に構築）

- Consumer（Go）

- MySQL（Docker Compose）

- バッチジョブ or スケジューラ（Go CLI）

## 主な機能

- ユーザーがお気に入りを押したイベントを Kafka へ送信

- Kafka から Consumer が非同期取得し MySQL に保存

- 定期バッチでランキング集計し更新

- DLQ（Dead Letter Queue）対応あり

## 技術スタック

- Go

- Kafka (Docker Compose)

- MySQL (Docker Compose)

- kafka-go ([segmentio](https://github.com/segmentio/kafka-go))
