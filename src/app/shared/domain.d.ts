declare global {
	type User = {
		id: number,
		givenName: string,
		familyName: string,
		email: string,
		updatedAt: Date,
		createdAt: Date
	}

	type Conversation = {
		id: number,
		name: string,
		users: User[],
		createdAt: Date,
		updatedAt: Date
	}
}

export {}