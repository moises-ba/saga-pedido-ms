package db

import (
	"context"

	"github.com/moises-ba/saga-pedido-ms/domain"
	"github.com/moises-ba/saga-pedido-ms/domain/entity"
	"github.com/moises-ba/saga-pedido-ms/domain/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type pedidoRepositoryMongo struct {
	dataBase   *mongo.Database
	collection *mongo.Collection
}

func NewPedidoMongoRepository(pDataBase *mongo.Database) repository.PedidoRepository {

	return &pedidoRepositoryMongo{dataBase: pDataBase, collection: pDataBase.Collection("pedido")}
}

func (r *pedidoRepositoryMongo) ListarPedidos(userId string) ([]*entity.Pedido, error) {
	return nil, nil
}

func (r *pedidoRepositoryMongo) ObterPedido(userId string, pPedidoId string) (*entity.Pedido, error) {

	pedidoEntity := &entity.Pedido{}
	singleResult := r.collection.FindOne(context.TODO(), bson.M{"_id": pPedidoId})

	if singleResult.Err() == mongo.ErrNoDocuments {
		logrus.Warnf("Pedido não encontrado para o id %v", pPedidoId)
		return nil, domain.ErrRecordNofFound
	}

	err := singleResult.Decode(pedidoEntity)
	if err != nil {
		return nil, err
	}

	if pedidoEntity.User.Id != userId {
		logrus.Warnf("Pedido não encontrado para o id %v e userId %v", pPedidoId, userId)
		return nil, domain.ErrRecordNofFound
	}

	return pedidoEntity, nil
}

func (r *pedidoRepositoryMongo) CriarPedido(pPedido *entity.Pedido) (*entity.Pedido, error) {

	_, err := r.collection.InsertOne(context.TODO(), pPedido)

	if err != nil && mongo.IsDuplicateKeyError(err) {

		err = domain.ErrDuplicateKey
	}

	return pPedido, err
}
