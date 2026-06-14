package repositories

import (
	"database/sql"

	"collab-code-platform/internal/models"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoomRepository(
	db *sql.DB,
) *RoomRepository {

	return &RoomRepository{
		db: db,
	}
}

func (r *RoomRepository) GetRoom(
	roomID string,
) (
	*models.Room,
	error,
) {

	query := `
		SELECT
			id,
			room_id,
			code,
			language,
			updated_at
		FROM rooms
		WHERE room_id = $1
	`

	room := &models.Room{}

	err := r.db.QueryRow(
		query,
		roomID,
	).Scan(
		&room.ID,
		&room.RoomID,
		&room.Code,
		&room.Language,
		&room.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r *RoomRepository) SaveRoom(
	roomID string,
	code string,
	language string,
) error {

	query := `
		INSERT INTO rooms (
			room_id,
			code,
			language
		)
		VALUES (
			$1,
			$2,
			$3
		)

		ON CONFLICT (room_id)

		DO UPDATE SET

		code = EXCLUDED.code,

		language = EXCLUDED.language,

		updated_at =
			CURRENT_TIMESTAMP
	`

	_, err := r.db.Exec(
		query,
		roomID,
		code,
		language,
	)

	return err
}
