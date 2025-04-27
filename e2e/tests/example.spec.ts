import { test, expect } from "@playwright/test";

test("ログインを要求される", async ({ page }) => {
  await page.goto("http://client:3000");
  await expect(page.locator("h1")).toHaveText("Please sign in");
});
