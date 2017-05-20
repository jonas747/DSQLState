package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordVoiceStates(t *testing.T) {
	t.Parallel()

	query := DiscordVoiceStates(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordVoiceStatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordVoiceState.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordVoiceStatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordVoiceStates(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordVoiceStatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordVoiceStateSlice{discordVoiceState}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordVoiceStatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordVoiceStateExists(tx, discordVoiceState.UserID, discordVoiceState.GuildID)
	if err != nil {
		t.Errorf("Unable to check if DiscordVoiceState exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordVoiceStateExistsG to return true, but got false.")
	}
}
func testDiscordVoiceStatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	discordVoiceStateFound, err := FindDiscordVoiceState(tx, discordVoiceState.UserID, discordVoiceState.GuildID)
	if err != nil {
		t.Error(err)
	}

	if discordVoiceStateFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordVoiceStatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordVoiceStates(tx).Bind(discordVoiceState); err != nil {
		t.Error(err)
	}
}

func testDiscordVoiceStatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordVoiceStates(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordVoiceStatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceStateOne := &DiscordVoiceState{}
	discordVoiceStateTwo := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceStateOne, discordVoiceStateDBTypes, false, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}
	if err = randomize.Struct(seed, discordVoiceStateTwo, discordVoiceStateDBTypes, false, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceStateOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordVoiceStateTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordVoiceStates(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordVoiceStatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordVoiceStateOne := &DiscordVoiceState{}
	discordVoiceStateTwo := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceStateOne, discordVoiceStateDBTypes, false, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}
	if err = randomize.Struct(seed, discordVoiceStateTwo, discordVoiceStateDBTypes, false, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceStateOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordVoiceStateTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordVoiceStatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordVoiceStatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx, discordVoiceStateColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordVoiceStateToOneDiscordChannelUsingChannel(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordVoiceState
	var foreign DiscordChannel

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordVoiceStateDBTypes, true, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ChannelID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Channel(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DiscordVoiceStateSlice{&local}
	if err = local.L.LoadChannel(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Channel == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Channel = nil
	if err = local.L.LoadChannel(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Channel == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDiscordVoiceStateToOneSetOpDiscordChannelUsingChannel(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordVoiceState
	var b, c DiscordChannel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordVoiceStateDBTypes, false, strmangle.SetComplement(discordVoiceStatePrimaryKeyColumns, discordVoiceStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, discordChannelDBTypes, false, strmangle.SetComplement(discordChannelPrimaryKeyColumns, discordChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, discordChannelDBTypes, false, strmangle.SetComplement(discordChannelPrimaryKeyColumns, discordChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DiscordChannel{&b, &c} {
		err = a.SetChannel(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Channel != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ChannelDiscordVoiceStates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ChannelID != x.ID {
			t.Error("foreign key was wrong value", a.ChannelID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ChannelID))
		reflect.Indirect(reflect.ValueOf(&a.ChannelID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ChannelID != x.ID {
			t.Error("foreign key was wrong value", a.ChannelID, x.ID)
		}
	}
}
func testDiscordVoiceStatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordVoiceState.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordVoiceStatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordVoiceStateSlice{discordVoiceState}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordVoiceStatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordVoiceStates(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordVoiceStateDBTypes = map[string]string{`ChannelID`: `bigint`, `Deaf`: `boolean`, `GuildID`: `bigint`, `Mute`: `boolean`, `SelfDeaf`: `boolean`, `SelfMute`: `boolean`, `SessionID`: `text`, `Surpress`: `boolean`, `UserID`: `bigint`}
	_                        = bytes.MinRead
)

func testDiscordVoiceStatesUpdate(t *testing.T) {
	t.Parallel()

	if len(discordVoiceStateColumns) == len(discordVoiceStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true, discordVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	if err = discordVoiceState.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordVoiceStatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordVoiceStateColumns) == len(discordVoiceStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordVoiceState := &DiscordVoiceState{}
	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordVoiceState, discordVoiceStateDBTypes, true, discordVoiceStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordVoiceStateColumns, discordVoiceStatePrimaryKeyColumns) {
		fields = discordVoiceStateColumns
	} else {
		fields = strmangle.SetComplement(
			discordVoiceStateColumns,
			discordVoiceStatePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordVoiceState))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordVoiceStateSlice{discordVoiceState}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordVoiceStatesUpsert(t *testing.T) {
	t.Parallel()

	if len(discordVoiceStateColumns) == len(discordVoiceStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordVoiceState := DiscordVoiceState{}
	if err = randomize.Struct(seed, &discordVoiceState, discordVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordVoiceState.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordVoiceState: %s", err)
	}

	count, err := DiscordVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordVoiceState, discordVoiceStateDBTypes, false, discordVoiceStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordVoiceState struct: %s", err)
	}

	if err = discordVoiceState.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordVoiceState: %s", err)
	}

	count, err = DiscordVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
