import unittest

from one_hot_encoder import fit_transform


class TestOneHotEncoder(unittest.TestCase):
    def test_(self):
        exp_transformed_cities = [
            ('Moscow', [0, 0, 1]),
            ('New York', [0, 1, 0]),
            ('Moscow', [0, 0, 1]),
            ('London', [1, 0, 0]),
        ]
        transformed_cities = fit_transform('Moscow', 'New York', 'Moscow', 'London')
        self.assertEqual(transformed_cities, exp_transformed_cities)

    def test_assert_NotIn(self):
        exp_transformed_cities = [
            ('Moscow', [0, 0, 1]),
            ('New York', [0, 1, 0]),
            ('Ankara', [0, 0, 1]),
            ('London', [1, 0, 0]),
        ]
        transformed_cities = fit_transform('Moscow', 'New York', 'Paris', 'London')
        self.assertNotIn(transformed_cities, exp_transformed_cities)

    def test_assert_exc(self):
        with self.assertRaises(TypeError):
            fit_transform('Moscow', 'New York', 'Paris', 'London', 'plant')

    def test_assertIsNotNone(self):
        exp_transformed_cities = [
        ('Moscow', [0, 0, 1]),
        ('New York', [0, 1, 0]),
        ('Moscow', [0, 0, 1]),
        ('London', [1, 0, 0]),
        ]
        self.assertIsNotNone(len(exp_transformed_cities))


if __name__ == '__main__':
    unittest.main()
