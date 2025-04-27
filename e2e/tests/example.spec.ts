import { test, expect } from "@playwright/test";

test("トップページが開ける", async ({ page }) => {
  await page.goto("http://app:3000");
  await expect(page).toHaveTitle(/Next\.js/);
});
