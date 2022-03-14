package uid

import "github.com/murtaza-udaipurwala/sme/db"

func NewUID(db *db.DB) (string, error) {
	var uid string
	var err error

	for {
		uid, err = genUID()
		if err != nil {
			return "", err
		}

		if uidExists(db, uid) {
			continue
		}

		break
	}

	return uid, nil
}
