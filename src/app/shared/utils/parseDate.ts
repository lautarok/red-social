export default function parseDate(dateStr: string) {
	const date = new Date(dateStr),
		nowDate = new Date()

	if (date.getTime() - nowDate.getTime() < 24*60*60*1000) {
		return `${
			date.getHours().toString().padStart(2, '0')
		}:${
			date.getMinutes().toString().padStart(2, '0')
		}`
	} else {
		return `${
			date.getDate().toString().padStart(2, '0')
		}/${
			date.getMonth().toString().padStart(2, '0')
		}/${
			date.getFullYear()
		} ${
			date.getHours().toString().padStart(2, '0')
		}:${
			date.getMinutes().toString().padStart(2, '0')
		}`
	}
}