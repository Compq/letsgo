import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import 'chapt4.dart';

void main() {
  runApp(App0402());
}

// class MyApp extends StatelessWidget {
//   @override
//   Widget build(BuildContext context) {
//     return MaterialApp(
//       title: 'Flutter Demo',
//       theme: ThemeData(
//         primarySwatch: Colors.blue,
//         visualDensity: VisualDensity.adaptivePlatformDensity,
//       ),
//       home: GooglePixel3XL1(),
//     );
//   }
// }

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
          appBar: AppBar(
            title: Text("MY 1srt Scaffold"),
            elevation: 100,
            brightness: Brightness.light,
          ),
          body: Center(
            child: Column(
              children: [
                Text(
                  "Hello",
                  textScaleFactor: 2.0,
                ),
                Text("Lonely for me")
              ],
            ),
          )
          // drawer: Drawer(
          //   child: Center(
          //     child: Text("I'm a drawer"),
          //   ),
          // ),
          ),
    );
  }
}
