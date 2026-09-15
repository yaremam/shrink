import 'package:flutter/material.dart';

void main() {
  runApp(const ShrinkApp());
}

class ShrinkApp extends StatelessWidget {
  const ShrinkApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Shrink!',
      theme: ThemeData(colorScheme: ColorScheme.fromSeed(seedColor: Colors.teal)),
      home: const Scaffold(
        body: Center(child: Text('Shrink!')),
      ),
    );
  }
}
