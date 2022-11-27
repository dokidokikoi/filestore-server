package init

import "file-store/internal/db/store/data"

func InitTask() {
	data.SetStoreDBFactory()
}
