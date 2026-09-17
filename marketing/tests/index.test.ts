import { experimental_AstroContainer as AstroContainer } from "astro/container";
import { describe, expect, it } from "vitest";
import Index from "../src/pages/index.astro";

describe("marketing homepage", () => {
	it("renders the Shrink! placeholder heading", async () => {
		const container = await AstroContainer.create();
		const result = await container.renderToString(Index);

		expect(result).toContain("<h1>Shrink!</h1>");
	});
});
