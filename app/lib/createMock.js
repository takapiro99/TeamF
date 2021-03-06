import { Server } from 'miragejs'
import { apiConfig } from '../config/api'

const endpoints = apiConfig.endpointsScheme

/**
 * APIモックを作成する関数を提供します。
 * @module createMock
 * @author Ritsuki KOKUBO
 * @see {@link https://miragejs.com/api/classes/association/}
 */

/**
 * カスタムヘッダーからトークンを取り出します
 * @param {Mirage.Request} request - Mirageのリクエストオブジェクト
 */
function getToken(request) {
  if (typeof request.requestHeaders !== 'undefined' && typeof request.requestHeaders.token !== 'undefined') {
    return request.requestHeaders.token
  } else {
    return undefined
  }
}

/**
 * APIのモックを作成し、起動します。
 * @returns {Mirage.Server} 作成したモックサーバー, 特に何もしなくても作成した時点でモックが有効になるようです
 */
export function createMock() {
  return new Server({
    routes() {
      this.urlPrefix = apiConfig.fqdn

      // セリフ一覧
      this.get(endpoints.getDialog, (_, request) => {
        const genre = request.queryParams.genre
        const offset = Number.parseInt(request.queryParams.offset)
        const limit = request.queryParams.limit
        let genreText
        if (genre == 'all') {
          genreText = ['anime', 'manga', 'book']
        } else {
          genreText = [genre]
        }
        const dialogs = []
        for (let i = 0; i < limit; i++) {
          const prefix = `-${offset + i + 1}-${genreText[i % genreText.length]}`
          dialogs.push({
            id: i + offset,
            content: `セリフ${prefix}`,
            title: `作品名${prefix}`,
            author: `著者${prefix}`,
            link: `リンク${prefix}`,
            style: `0`,
            source: `著作権${prefix}`
          })
        }
        return {
          message: 'ok',
          schema: dialogs
        }
      })

      // セリフ詳細
      this.get(endpoints.getDialogDetail, (_, request) => {
        const id = request.params.id
        const offset = Number.parseInt(request.queryParams.offset)
        const limit = Number.parseInt(request.queryParams.limit)
        const prefix = `-${id}`
        const comments = []
        for (let i = 0; i < limit; i++) {
          comments.push({
            content: `セリフ${prefix}に対するコメント${i}`,
            user: {
              id: i,
              firebase_uid: i,
              display_name: `ユーザ${i}`,
              photo_url: 'author.png'
            }
          })
        }
        return {
          message: 'ok',
          dialog: {
            id: id,
            content: `セリフ${prefix}`,
            title: `作品名${prefix}`,
            author: `著者${prefix}`,
            link: `リンク${prefix}`,
            style: `0`,
            source: `著作権${prefix}`
          },
          comments
        }
      })

      // セリフ投稿
      this.post(endpoints.postDialog, (_, request) => {
        // ...
        const token = getToken(request)
        const dialog = JSON.parse(request.requestBody)
        console.info('On Mock API: ')
        console.log(dialog)
        return {
          message: 'ok'
        }
      })

      // コメント投稿
      this.post(endpoints.postComment, (_, request) => {
        // ...
        const token = getToken(request)
        const id = request.params.id
        const comment = JSON.parse(request.requestBody)
        console.info('On Mock API: ')
        console.log(comment)
        return {
          message: 'ok'
        }
      })
    }
  })
}
