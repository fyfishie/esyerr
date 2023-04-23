/*
 * @Author: fyfishie
 * @Date: 2023-04-23:15
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-04-23:15
 * @@email: fyfishie@outlook.com
 * @Description: :)
 */
package esyerr

func AutoPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
