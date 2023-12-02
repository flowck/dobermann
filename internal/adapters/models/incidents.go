// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Incident is an object representing the database table.
type Incident struct {
	ID         string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	MonitorID  string    `boil:"monitor_id" json:"monitor_id" toml:"monitor_id" yaml:"monitor_id"`
	IsResolved bool      `boil:"is_resolved" json:"is_resolved" toml:"is_resolved" yaml:"is_resolved"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *incidentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L incidentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var IncidentColumns = struct {
	ID         string
	MonitorID  string
	IsResolved string
	CreatedAt  string
}{
	ID:         "id",
	MonitorID:  "monitor_id",
	IsResolved: "is_resolved",
	CreatedAt:  "created_at",
}

var IncidentTableColumns = struct {
	ID         string
	MonitorID  string
	IsResolved string
	CreatedAt  string
}{
	ID:         "incidents.id",
	MonitorID:  "incidents.monitor_id",
	IsResolved: "incidents.is_resolved",
	CreatedAt:  "incidents.created_at",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var IncidentWhere = struct {
	ID         whereHelperstring
	MonitorID  whereHelperstring
	IsResolved whereHelperbool
	CreatedAt  whereHelpertime_Time
}{
	ID:         whereHelperstring{field: "\"incidents\".\"id\""},
	MonitorID:  whereHelperstring{field: "\"incidents\".\"monitor_id\""},
	IsResolved: whereHelperbool{field: "\"incidents\".\"is_resolved\""},
	CreatedAt:  whereHelpertime_Time{field: "\"incidents\".\"created_at\""},
}

// IncidentRels is where relationship names are stored.
var IncidentRels = struct {
	Monitor         string
	IncidentActions string
}{
	Monitor:         "Monitor",
	IncidentActions: "IncidentActions",
}

// incidentR is where relationships are stored.
type incidentR struct {
	Monitor         *Monitor            `boil:"Monitor" json:"Monitor" toml:"Monitor" yaml:"Monitor"`
	IncidentActions IncidentActionSlice `boil:"IncidentActions" json:"IncidentActions" toml:"IncidentActions" yaml:"IncidentActions"`
}

// NewStruct creates a new relationship struct
func (*incidentR) NewStruct() *incidentR {
	return &incidentR{}
}

func (r *incidentR) GetMonitor() *Monitor {
	if r == nil {
		return nil
	}
	return r.Monitor
}

func (r *incidentR) GetIncidentActions() IncidentActionSlice {
	if r == nil {
		return nil
	}
	return r.IncidentActions
}

// incidentL is where Load methods for each relationship are stored.
type incidentL struct{}

var (
	incidentAllColumns            = []string{"id", "monitor_id", "is_resolved", "created_at"}
	incidentColumnsWithoutDefault = []string{"id", "monitor_id"}
	incidentColumnsWithDefault    = []string{"is_resolved", "created_at"}
	incidentPrimaryKeyColumns     = []string{"id"}
	incidentGeneratedColumns      = []string{}
)

type (
	// IncidentSlice is an alias for a slice of pointers to Incident.
	// This should almost always be used instead of []Incident.
	IncidentSlice []*Incident
	// IncidentHook is the signature for custom Incident hook methods
	IncidentHook func(context.Context, boil.ContextExecutor, *Incident) error

	incidentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	incidentType                 = reflect.TypeOf(&Incident{})
	incidentMapping              = queries.MakeStructMapping(incidentType)
	incidentPrimaryKeyMapping, _ = queries.BindMapping(incidentType, incidentMapping, incidentPrimaryKeyColumns)
	incidentInsertCacheMut       sync.RWMutex
	incidentInsertCache          = make(map[string]insertCache)
	incidentUpdateCacheMut       sync.RWMutex
	incidentUpdateCache          = make(map[string]updateCache)
	incidentUpsertCacheMut       sync.RWMutex
	incidentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var incidentAfterSelectHooks []IncidentHook

var incidentBeforeInsertHooks []IncidentHook
var incidentAfterInsertHooks []IncidentHook

var incidentBeforeUpdateHooks []IncidentHook
var incidentAfterUpdateHooks []IncidentHook

var incidentBeforeDeleteHooks []IncidentHook
var incidentAfterDeleteHooks []IncidentHook

var incidentBeforeUpsertHooks []IncidentHook
var incidentAfterUpsertHooks []IncidentHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Incident) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range incidentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Incident) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range incidentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Incident) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range incidentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Incident) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range incidentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Incident) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range incidentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Incident) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range incidentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Incident) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range incidentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Incident) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range incidentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Incident) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range incidentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddIncidentHook registers your hook function for all future operations.
func AddIncidentHook(hookPoint boil.HookPoint, incidentHook IncidentHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		incidentAfterSelectHooks = append(incidentAfterSelectHooks, incidentHook)
	case boil.BeforeInsertHook:
		incidentBeforeInsertHooks = append(incidentBeforeInsertHooks, incidentHook)
	case boil.AfterInsertHook:
		incidentAfterInsertHooks = append(incidentAfterInsertHooks, incidentHook)
	case boil.BeforeUpdateHook:
		incidentBeforeUpdateHooks = append(incidentBeforeUpdateHooks, incidentHook)
	case boil.AfterUpdateHook:
		incidentAfterUpdateHooks = append(incidentAfterUpdateHooks, incidentHook)
	case boil.BeforeDeleteHook:
		incidentBeforeDeleteHooks = append(incidentBeforeDeleteHooks, incidentHook)
	case boil.AfterDeleteHook:
		incidentAfterDeleteHooks = append(incidentAfterDeleteHooks, incidentHook)
	case boil.BeforeUpsertHook:
		incidentBeforeUpsertHooks = append(incidentBeforeUpsertHooks, incidentHook)
	case boil.AfterUpsertHook:
		incidentAfterUpsertHooks = append(incidentAfterUpsertHooks, incidentHook)
	}
}

// One returns a single incident record from the query.
func (q incidentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Incident, error) {
	o := &Incident{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for incidents")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Incident records from the query.
func (q incidentQuery) All(ctx context.Context, exec boil.ContextExecutor) (IncidentSlice, error) {
	var o []*Incident

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Incident slice")
	}

	if len(incidentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Incident records in the query.
func (q incidentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count incidents rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q incidentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if incidents exists")
	}

	return count > 0, nil
}

// Monitor pointed to by the foreign key.
func (o *Incident) Monitor(mods ...qm.QueryMod) monitorQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.MonitorID),
	}

	queryMods = append(queryMods, mods...)

	return Monitors(queryMods...)
}

// IncidentActions retrieves all the incident_action's IncidentActions with an executor.
func (o *Incident) IncidentActions(mods ...qm.QueryMod) incidentActionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"incident_actions\".\"incident_id\"=?", o.ID),
	)

	return IncidentActions(queryMods...)
}

// LoadMonitor allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (incidentL) LoadMonitor(ctx context.Context, e boil.ContextExecutor, singular bool, maybeIncident interface{}, mods queries.Applicator) error {
	var slice []*Incident
	var object *Incident

	if singular {
		var ok bool
		object, ok = maybeIncident.(*Incident)
		if !ok {
			object = new(Incident)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeIncident)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeIncident))
			}
		}
	} else {
		s, ok := maybeIncident.(*[]*Incident)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeIncident)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeIncident))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &incidentR{}
		}
		args = append(args, object.MonitorID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &incidentR{}
			}

			for _, a := range args {
				if a == obj.MonitorID {
					continue Outer
				}
			}

			args = append(args, obj.MonitorID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`monitors`),
		qm.WhereIn(`monitors.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Monitor")
	}

	var resultSlice []*Monitor
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Monitor")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for monitors")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for monitors")
	}

	if len(monitorAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Monitor = foreign
		if foreign.R == nil {
			foreign.R = &monitorR{}
		}
		foreign.R.Incidents = append(foreign.R.Incidents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MonitorID == foreign.ID {
				local.R.Monitor = foreign
				if foreign.R == nil {
					foreign.R = &monitorR{}
				}
				foreign.R.Incidents = append(foreign.R.Incidents, local)
				break
			}
		}
	}

	return nil
}

// LoadIncidentActions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (incidentL) LoadIncidentActions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeIncident interface{}, mods queries.Applicator) error {
	var slice []*Incident
	var object *Incident

	if singular {
		var ok bool
		object, ok = maybeIncident.(*Incident)
		if !ok {
			object = new(Incident)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeIncident)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeIncident))
			}
		}
	} else {
		s, ok := maybeIncident.(*[]*Incident)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeIncident)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeIncident))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &incidentR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &incidentR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`incident_actions`),
		qm.WhereIn(`incident_actions.incident_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load incident_actions")
	}

	var resultSlice []*IncidentAction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice incident_actions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on incident_actions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for incident_actions")
	}

	if len(incidentActionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.IncidentActions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &incidentActionR{}
			}
			foreign.R.Incident = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.IncidentID {
				local.R.IncidentActions = append(local.R.IncidentActions, foreign)
				if foreign.R == nil {
					foreign.R = &incidentActionR{}
				}
				foreign.R.Incident = local
				break
			}
		}
	}

	return nil
}

// SetMonitor of the incident to the related item.
// Sets o.R.Monitor to related.
// Adds o to related.R.Incidents.
func (o *Incident) SetMonitor(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Monitor) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"incidents\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"monitor_id"}),
		strmangle.WhereClause("\"", "\"", 2, incidentPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MonitorID = related.ID
	if o.R == nil {
		o.R = &incidentR{
			Monitor: related,
		}
	} else {
		o.R.Monitor = related
	}

	if related.R == nil {
		related.R = &monitorR{
			Incidents: IncidentSlice{o},
		}
	} else {
		related.R.Incidents = append(related.R.Incidents, o)
	}

	return nil
}

// AddIncidentActions adds the given related objects to the existing relationships
// of the incident, optionally inserting them as new records.
// Appends related to o.R.IncidentActions.
// Sets related.R.Incident appropriately.
func (o *Incident) AddIncidentActions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*IncidentAction) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.IncidentID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"incident_actions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"incident_id"}),
				strmangle.WhereClause("\"", "\"", 2, incidentActionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.IncidentID = o.ID
		}
	}

	if o.R == nil {
		o.R = &incidentR{
			IncidentActions: related,
		}
	} else {
		o.R.IncidentActions = append(o.R.IncidentActions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &incidentActionR{
				Incident: o,
			}
		} else {
			rel.R.Incident = o
		}
	}
	return nil
}

// Incidents retrieves all the records using an executor.
func Incidents(mods ...qm.QueryMod) incidentQuery {
	mods = append(mods, qm.From("\"incidents\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"incidents\".*"})
	}

	return incidentQuery{q}
}

// FindIncident retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindIncident(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Incident, error) {
	incidentObj := &Incident{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"incidents\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, incidentObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from incidents")
	}

	if err = incidentObj.doAfterSelectHooks(ctx, exec); err != nil {
		return incidentObj, err
	}

	return incidentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Incident) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no incidents provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(incidentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	incidentInsertCacheMut.RLock()
	cache, cached := incidentInsertCache[key]
	incidentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			incidentAllColumns,
			incidentColumnsWithDefault,
			incidentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(incidentType, incidentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(incidentType, incidentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"incidents\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"incidents\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into incidents")
	}

	if !cached {
		incidentInsertCacheMut.Lock()
		incidentInsertCache[key] = cache
		incidentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Incident.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Incident) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	incidentUpdateCacheMut.RLock()
	cache, cached := incidentUpdateCache[key]
	incidentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			incidentAllColumns,
			incidentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update incidents, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"incidents\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, incidentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(incidentType, incidentMapping, append(wl, incidentPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update incidents row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for incidents")
	}

	if !cached {
		incidentUpdateCacheMut.Lock()
		incidentUpdateCache[key] = cache
		incidentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q incidentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for incidents")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for incidents")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o IncidentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), incidentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"incidents\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, incidentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in incident slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all incident")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Incident) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no incidents provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(incidentColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	incidentUpsertCacheMut.RLock()
	cache, cached := incidentUpsertCache[key]
	incidentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			incidentAllColumns,
			incidentColumnsWithDefault,
			incidentColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			incidentAllColumns,
			incidentPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert incidents, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(incidentPrimaryKeyColumns))
			copy(conflict, incidentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"incidents\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(incidentType, incidentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(incidentType, incidentMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert incidents")
	}

	if !cached {
		incidentUpsertCacheMut.Lock()
		incidentUpsertCache[key] = cache
		incidentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Incident record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Incident) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Incident provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), incidentPrimaryKeyMapping)
	sql := "DELETE FROM \"incidents\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from incidents")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for incidents")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q incidentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no incidentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from incidents")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for incidents")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o IncidentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(incidentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), incidentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"incidents\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, incidentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from incident slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for incidents")
	}

	if len(incidentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Incident) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindIncident(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IncidentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := IncidentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), incidentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"incidents\".* FROM \"incidents\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, incidentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in IncidentSlice")
	}

	*o = slice

	return nil
}

// IncidentExists checks if the Incident row exists.
func IncidentExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"incidents\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if incidents exists")
	}

	return exists, nil
}

// Exists checks if the Incident row exists.
func (o *Incident) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return IncidentExists(ctx, exec, o.ID)
}
