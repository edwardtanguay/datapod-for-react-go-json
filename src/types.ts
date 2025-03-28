import { z } from "zod";

export const FlashcardSchema = z.object({
	suuid: z
		.string()
		.length(6, "suuid must be exactly 6 characters long")
		.regex(/^[A-Za-z0-9]+$/, "suuid can only contain uppercase/lowercase letters and numbers"),
	category: z.string(),
	front: z.string(),
	back: z.string(),
});

export const FrontendFlashcardSchema = FlashcardSchema.extend({
	isOpen: z.boolean(),
});

export type Flashcard = z.infer<typeof FlashcardSchema>;
export type FrontendFlashcard = z.infer<typeof FrontendFlashcardSchema>;