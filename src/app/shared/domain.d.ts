declare global {
	type User = {
		id: number,
		givenName: string,
		familyName: string,
		email: string,
		updatedAt: string,
		createdAt: string
	}

	type Message = {
		id: number,
		text: string,
		fromUserId: number,
		conversationId: number,
		createdAt: string,
		updatedAt: string
	}

	type Conversation = {
		id: number,
		name: string,
		users: User[],
		messages: Message[],
		lastMessage?: Message,
		createdAt: string,
		updatedAt: string
	}
}

export {}