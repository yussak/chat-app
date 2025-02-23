import NextAuth from "next-auth";
import GoogleProvider from "next-auth/providers/google";

const handler = NextAuth({
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_ID ?? "",
      clientSecret: process.env.GOOGLE_SECRET ?? "",
    }),
  ],
  callbacks: {
    async signIn({ user }) {
      try {
        const response = await fetch(`http://server:8080/users/signin`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            email: user.email,
            name: user.name,
            image: user.image,
          }),
        });

        const dbUser = await response.json();
        // セッションにDBのユーザーIDを保存するため
        user.id = dbUser.ID;
        return true;
      } catch (error) {
        console.error("ユーザー情報送信失敗", error);
        return false;
      }
    },
    async jwt({ token, user }) {
      // サインイン時にユーザーIDをトークンに追加
      if (user) {
        token.id = user.id;
      }
      return token;
    },
    async session({ session, token }) {
      // トークンからユーザーIDをセッションに追加
      if (session.user) {
        session.user.id = token.id;
      }
      return session;
    },
  },
});

export { handler as GET, handler as POST };
