package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	util "github.com/thewackyindian/3iOj/utils"
)

func createRandomContest(t *testing.T) Contest {
	cur := time.Now();
	arg := CreateContestParams{
		ContestName:    util.RandomContestName(),
		StartTime: 		cur.Add(1 * time.Hour),
		EndTime:       	cur.Add(3 * time.Hour),
		RegistrationStart:cur,
		RegistrationEnd: cur.Add(55 * time.Minute),
		//here we will call for two blogs creation and the link their IDs
		AnnouncementBlog: util.RandomInt(100, 200),
		EditorialBlog: util.RandomInt(100, 200),
	}

	contest, err := testQueries.CreateContest(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, contest)

	require.Equal(t, arg.ContestName, contest.ContestName)
	require.Equal(t, arg.StartTime, contest.StartTime)
	require.Equal(t, arg.EndTime, contest.EndTime)
	require.Equal(t, arg.RegistrationStart, contest.RegistrationStart)
	require.Equal(t, arg.RegistrationEnd, contest.RegistrationEnd)
	require.NotZero(t, contest.CreatedAt)

	return contest
}

func TestCreateContest(t *testing.T) {
	createRandomContest(t)
}

func TestGetContest(t *testing.T) {
	contest1 := createRandomContest(t)
	contest2, err := testQueries.GetContest(context.Background(), contest1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, contest2)
	require.Equal(t, contest1.ContestName, contest2.ContestName)
	require.Equal(t, contest1.StartTime, contest2.StartTime)
	require.Equal(t, contest1.EndTime, contest2.EndTime)
	require.Equal(t, contest1.RegistrationStart, contest2.RegistrationStart)
	require.Equal(t, contest1.RegistrationEnd, contest2.RegistrationEnd)
	require.WithinDuration(t, contest1.CreatedAt, contest2.CreatedAt, time.Second)
}
func TestDeleteContest(t *testing.T) {
	contest1 := createRandomContest(t)
	err := testQueries.DeleteContest(context.Background(), contest1.ID)
	require.NoError(t, err)

	contest2, err := testQueries.GetContest(context.Background(), contest1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, contest2)
}