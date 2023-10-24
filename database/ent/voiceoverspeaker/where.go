// Code generated by ent, DO NOT EDIT.

package voiceoverspeaker

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLTE(FieldID, id))
}

// NameInWorker applies equality check predicate on the "name_in_worker" field. It's identical to NameInWorkerEQ.
func NameInWorker(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldNameInWorker, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldName, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldIsActive, v))
}

// IsDefault applies equality check predicate on the "is_default" field. It's identical to IsDefaultEQ.
func IsDefault(v bool) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldIsDefault, v))
}

// IsHidden applies equality check predicate on the "is_hidden" field. It's identical to IsHiddenEQ.
func IsHidden(v bool) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldIsHidden, v))
}

// Locale applies equality check predicate on the "locale" field. It's identical to LocaleEQ.
func Locale(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldLocale, v))
}

// ModelID applies equality check predicate on the "model_id" field. It's identical to ModelIDEQ.
func ModelID(v uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldModelID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameInWorkerEQ applies the EQ predicate on the "name_in_worker" field.
func NameInWorkerEQ(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldNameInWorker, v))
}

// NameInWorkerNEQ applies the NEQ predicate on the "name_in_worker" field.
func NameInWorkerNEQ(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNEQ(FieldNameInWorker, v))
}

// NameInWorkerIn applies the In predicate on the "name_in_worker" field.
func NameInWorkerIn(vs ...string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldIn(FieldNameInWorker, vs...))
}

// NameInWorkerNotIn applies the NotIn predicate on the "name_in_worker" field.
func NameInWorkerNotIn(vs ...string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNotIn(FieldNameInWorker, vs...))
}

// NameInWorkerGT applies the GT predicate on the "name_in_worker" field.
func NameInWorkerGT(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGT(FieldNameInWorker, v))
}

// NameInWorkerGTE applies the GTE predicate on the "name_in_worker" field.
func NameInWorkerGTE(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGTE(FieldNameInWorker, v))
}

// NameInWorkerLT applies the LT predicate on the "name_in_worker" field.
func NameInWorkerLT(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLT(FieldNameInWorker, v))
}

// NameInWorkerLTE applies the LTE predicate on the "name_in_worker" field.
func NameInWorkerLTE(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLTE(FieldNameInWorker, v))
}

// NameInWorkerContains applies the Contains predicate on the "name_in_worker" field.
func NameInWorkerContains(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldContains(FieldNameInWorker, v))
}

// NameInWorkerHasPrefix applies the HasPrefix predicate on the "name_in_worker" field.
func NameInWorkerHasPrefix(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldHasPrefix(FieldNameInWorker, v))
}

// NameInWorkerHasSuffix applies the HasSuffix predicate on the "name_in_worker" field.
func NameInWorkerHasSuffix(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldHasSuffix(FieldNameInWorker, v))
}

// NameInWorkerEqualFold applies the EqualFold predicate on the "name_in_worker" field.
func NameInWorkerEqualFold(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEqualFold(FieldNameInWorker, v))
}

// NameInWorkerContainsFold applies the ContainsFold predicate on the "name_in_worker" field.
func NameInWorkerContainsFold(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldContainsFold(FieldNameInWorker, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldContainsFold(FieldName, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNEQ(FieldIsActive, v))
}

// IsDefaultEQ applies the EQ predicate on the "is_default" field.
func IsDefaultEQ(v bool) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldIsDefault, v))
}

// IsDefaultNEQ applies the NEQ predicate on the "is_default" field.
func IsDefaultNEQ(v bool) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNEQ(FieldIsDefault, v))
}

// IsHiddenEQ applies the EQ predicate on the "is_hidden" field.
func IsHiddenEQ(v bool) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldIsHidden, v))
}

// IsHiddenNEQ applies the NEQ predicate on the "is_hidden" field.
func IsHiddenNEQ(v bool) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNEQ(FieldIsHidden, v))
}

// LocaleEQ applies the EQ predicate on the "locale" field.
func LocaleEQ(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldLocale, v))
}

// LocaleNEQ applies the NEQ predicate on the "locale" field.
func LocaleNEQ(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNEQ(FieldLocale, v))
}

// LocaleIn applies the In predicate on the "locale" field.
func LocaleIn(vs ...string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldIn(FieldLocale, vs...))
}

// LocaleNotIn applies the NotIn predicate on the "locale" field.
func LocaleNotIn(vs ...string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNotIn(FieldLocale, vs...))
}

// LocaleGT applies the GT predicate on the "locale" field.
func LocaleGT(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGT(FieldLocale, v))
}

// LocaleGTE applies the GTE predicate on the "locale" field.
func LocaleGTE(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGTE(FieldLocale, v))
}

// LocaleLT applies the LT predicate on the "locale" field.
func LocaleLT(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLT(FieldLocale, v))
}

// LocaleLTE applies the LTE predicate on the "locale" field.
func LocaleLTE(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLTE(FieldLocale, v))
}

// LocaleContains applies the Contains predicate on the "locale" field.
func LocaleContains(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldContains(FieldLocale, v))
}

// LocaleHasPrefix applies the HasPrefix predicate on the "locale" field.
func LocaleHasPrefix(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldHasPrefix(FieldLocale, v))
}

// LocaleHasSuffix applies the HasSuffix predicate on the "locale" field.
func LocaleHasSuffix(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldHasSuffix(FieldLocale, v))
}

// LocaleEqualFold applies the EqualFold predicate on the "locale" field.
func LocaleEqualFold(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEqualFold(FieldLocale, v))
}

// LocaleContainsFold applies the ContainsFold predicate on the "locale" field.
func LocaleContainsFold(v string) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldContainsFold(FieldLocale, v))
}

// ModelIDEQ applies the EQ predicate on the "model_id" field.
func ModelIDEQ(v uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldModelID, v))
}

// ModelIDNEQ applies the NEQ predicate on the "model_id" field.
func ModelIDNEQ(v uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNEQ(FieldModelID, v))
}

// ModelIDIn applies the In predicate on the "model_id" field.
func ModelIDIn(vs ...uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldIn(FieldModelID, vs...))
}

// ModelIDNotIn applies the NotIn predicate on the "model_id" field.
func ModelIDNotIn(vs ...uuid.UUID) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNotIn(FieldModelID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasVoiceovers applies the HasEdge predicate on the "voiceovers" edge.
func HasVoiceovers() predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VoiceoversTable, VoiceoversColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVoiceoversWith applies the HasEdge predicate on the "voiceovers" edge with a given conditions (other predicates).
func HasVoiceoversWith(preds ...predicate.Voiceover) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VoiceoversInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VoiceoversTable, VoiceoversColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVoiceoverModels applies the HasEdge predicate on the "voiceover_models" edge.
func HasVoiceoverModels() predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VoiceoverModelsTable, VoiceoverModelsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVoiceoverModelsWith applies the HasEdge predicate on the "voiceover_models" edge with a given conditions (other predicates).
func HasVoiceoverModelsWith(preds ...predicate.VoiceoverModel) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VoiceoverModelsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VoiceoverModelsTable, VoiceoverModelsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VoiceoverSpeaker) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VoiceoverSpeaker) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VoiceoverSpeaker) predicate.VoiceoverSpeaker {
	return predicate.VoiceoverSpeaker(func(s *sql.Selector) {
		p(s.Not())
	})
}
