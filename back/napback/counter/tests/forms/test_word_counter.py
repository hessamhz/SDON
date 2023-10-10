from counter.forms import WordCountForm
from django.test import TestCase
from faker import Faker


class WordCountFormTest(TestCase):
    """
    Test case for the WordCountForm class.
    """

    def setUp(self):
        self.faker = Faker()

    def test_form_with_valid_text(self):
        """
        Test the form with valid text input.
        The form should be considered valid and the cleaned data should match the input.
        """
        fake_text = self.faker.paragraph()
        form_data = {"text": fake_text}
        form = WordCountForm(data=form_data)
        self.assertTrue(form.is_valid())
        self.assertEqual(form.cleaned_data.get("text"), fake_text)

    def test_form_with_empty_text(self):
        """
        Test the form with an empty text input.
        The form should be considered invalid and an appropriate error message should be displayed.
        """
        form_data = {"text": ""}
        form = WordCountForm(data=form_data)
        self.assertFalse(form.is_valid())
        self.assertIn("text", form.errors)
        self.assertEqual(
            form.errors.get("text"), ["Text field should not be empty."]
        )

    def test_form_with_whitespace_text(self):
        """
        Test the form with only whitespace input.
        The form should be considered invalid and an appropriate error message should be displayed.
        """
        form_data = {"text": "     "}
        form = WordCountForm(data=form_data)
        self.assertFalse(form.is_valid())
        self.assertIn("text", form.errors)
        self.assertEqual(
            form.errors.get("text"), ["Text field should not be empty."]
        )

    def test_form_with_strip_whitespace(self):
        """
        Test the form with text input containing leading and trailing whitespace.
        The form should be considered valid after stripping whitespace, and the cleaned data should match the stripped input.
        """
        fake_text = "    Some text with leading and trailing whitespace.   "
        form_data = {"text": fake_text}
        form = WordCountForm(data=form_data)
        self.assertTrue(form.is_valid())
        self.assertEqual(
            form.cleaned_data.get("text"),
            fake_text.strip(),
        )
