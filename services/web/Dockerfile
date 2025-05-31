FROM node:22-alpine AS build

ENV NEXT_TELEMETRY_DISABLED=1
WORKDIR /app

COPY ./services/web /app

RUN apk add --no-cache libc6-compat
RUN corepack enable yarn
RUN yarn install
RUN yarn build


FROM node:22-alpine

ENV NEXT_TELEMETRY_DISABLED=1
ENV NODE_ENV=production
ENV PORT=8080
WORKDIR /app

COPY --from=build /app/public ./public
COPY --from=build /app/.next/standalone ./
COPY --from=build /app/.next/static ./.next/static

EXPOSE 8080
CMD ["node", "server.js"]