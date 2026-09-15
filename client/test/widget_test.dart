import 'package:flutter_test/flutter_test.dart';

import 'package:shrink/main.dart';

void main() {
  testWidgets('shows the Shrink! placeholder', (WidgetTester tester) async {
    await tester.pumpWidget(const ShrinkApp());

    expect(find.text('Shrink!'), findsOneWidget);
  });
}
