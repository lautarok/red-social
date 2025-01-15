declare global {
	type User = {
		id: number,
		givenName: string,
		familyName: string,
		email: string,
		updatedAt: Date,
		createdAt: Date
	}

	type Message = {
		id: number,
		text: string,
		fromUserId: number,
		createdAt: Date,
		updatedAt: Date
	}

	type Conversation = {
		id: number,
		name: string,
		users: User[],
		messages: Message[],
		createdAt: Date,
		updatedAt: Date
	}
}

export {}