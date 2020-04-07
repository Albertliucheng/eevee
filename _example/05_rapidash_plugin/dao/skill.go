package dao

import (
	"context"
	"rapidashplugin/entity"

	"go.knocknote.io/rapidash"
	"golang.org/x/xerrors"
)

type Skill interface {
	Count(context.Context) (int64, error)
	Create(context.Context, *entity.Skill) error
	Delete(context.Context, *entity.Skill) error
	DeleteByID(context.Context, uint64) error
	DeleteByIDs(context.Context, []uint64) error
	FindAll(context.Context) (entity.Skills, error)
	FindByID(context.Context, uint64) (*entity.Skill, error)
	FindByIDs(context.Context, []uint64) (entity.Skills, error)
	Update(context.Context, *entity.Skill) error
	UpdateByID(context.Context, uint64, map[string]interface{}) error
	UpdateByIDs(context.Context, []uint64, map[string]interface{}) error
}

type SkillImpl struct {
	tx *rapidash.Tx
}

func NewSkill(ctx context.Context, tx *rapidash.Tx) Skill {
	return &SkillImpl{tx: tx}
}

// generated by eevee
func (d *SkillImpl) Count(ctx context.Context) (r int64, e error) {
	builder := rapidash.NewQueryBuilder("skills")
	count, err := d.tx.CountByQueryBuilderContext(ctx, builder)
	if err != nil {
		return 0, xerrors.Errorf("failed to Count: %w", err)
	}
	return int64(count), nil
}

// generated by eevee
func (d *SkillImpl) Create(ctx context.Context, value *entity.Skill) (e error) {
	id, err := d.tx.CreateByTableContext(ctx, "skills", value)
	if err != nil {
		return xerrors.Errorf("failed to Create: %w", err)
	}
	value.ID = uint64(id)
	return nil
}

// generated by eevee
func (d *SkillImpl) Delete(ctx context.Context, value *entity.Skill) (e error) {
	builder := rapidash.NewQueryBuilder("skills").Eq("id", value.ID)
	if err := d.tx.DeleteByQueryBuilderContext(ctx, builder); err != nil {
		return xerrors.Errorf("failed to Delete: %w", err)
	}
	return nil
}

// generated by eevee
func (d *SkillImpl) DeleteByID(ctx context.Context, a0 uint64) (e error) {
	builder := rapidash.NewQueryBuilder("skills").Eq("id", a0)
	if err := d.tx.DeleteByQueryBuilderContext(ctx, builder); err != nil {
		return xerrors.Errorf("failed to Delete: %w", err)
	}
	return nil
}

// generated by eevee
func (d *SkillImpl) DeleteByIDs(ctx context.Context, a0 []uint64) (e error) {
	builder := rapidash.NewQueryBuilder("skills").In("id", a0)
	if err := d.tx.DeleteByQueryBuilderContext(ctx, builder); err != nil {
		return xerrors.Errorf("failed to Delete: %w", err)
	}
	return nil
}

// generated by eevee
func (d *SkillImpl) FindAll(ctx context.Context) (r entity.Skills, e error) {
	values := entity.Skills{}
	if err := d.tx.FindAllByTable("skills", &values); err != nil {
		return values, xerrors.Errorf("failed to FindAll %w", err)
	}
	return values, nil
}

// generated by eevee
func (d *SkillImpl) FindByID(ctx context.Context, a0 uint64) (r *entity.Skill, e error) {
	builder := rapidash.NewQueryBuilder("skills").Eq("id", a0)
	var value entity.Skill
	if err := d.tx.FindByQueryBuilderContext(ctx, builder, &value); err != nil {
		return nil, xerrors.Errorf("failed to Find: %w", err)
	}
	return &value, nil
}

// generated by eevee
func (d *SkillImpl) FindByIDs(ctx context.Context, a0 []uint64) (r entity.Skills, e error) {
	builder := rapidash.NewQueryBuilder("skills").In("id", a0)
	values := entity.Skills{}
	if err := d.tx.FindByQueryBuilderContext(ctx, builder, &values); err != nil {
		return nil, xerrors.Errorf("failed to Find: %w", err)
	}
	return values, nil
}

// generated by eevee
func (d *SkillImpl) Update(ctx context.Context, value *entity.Skill) (e error) {
	updateMap := map[string]interface{}{"skill_effect": value.SkillEffect}
	builder := rapidash.NewQueryBuilder("skills").Eq("id", value.ID)
	if err := d.tx.UpdateByQueryBuilderContext(ctx, builder, updateMap); err != nil {
		return xerrors.Errorf("failed to Update: %w", err)
	}
	return nil
}

// generated by eevee
func (d *SkillImpl) UpdateByID(ctx context.Context, a0 uint64, updateMap map[string]interface{}) (e error) {
	builder := rapidash.NewQueryBuilder("skills").Eq("id", a0)
	if err := d.tx.UpdateByQueryBuilderContext(ctx, builder, updateMap); err != nil {
		return xerrors.Errorf("failed to Update: %w", err)
	}
	return nil
}

// generated by eevee
func (d *SkillImpl) UpdateByIDs(ctx context.Context, a0 []uint64, updateMap map[string]interface{}) (e error) {
	builder := rapidash.NewQueryBuilder("skills").In("id", a0)
	if err := d.tx.UpdateByQueryBuilderContext(ctx, builder, updateMap); err != nil {
		return xerrors.Errorf("failed to Update: %w", err)
	}
	return nil
}