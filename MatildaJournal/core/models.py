from django.db import models

class Publication(models.Model):
    title = models.CharField('title', max_length=100)
    description = models.CharField('description', max_length=300)
    photo = models.ImageField(null = True, blank= True)

    def __str__(self):
        return self.title
