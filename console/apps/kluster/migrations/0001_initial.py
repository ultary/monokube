# Generated by Django 5.0.7 on 2024-07-14 06:07

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='ResourceStatus',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('api_group', models.CharField(blank=True, default='', max_length=32)),
                ('api_version', models.CharField(max_length=16)),
                ('kind', models.CharField(max_length=128)),
                ('name', models.CharField(max_length=253)),
                ('namespace', models.CharField(blank=True, default='', max_length=63)),
                ('requested', models.TextField(max_length=1048576)),
                ('status', models.JSONField(max_length=2097152)),
                ('resource_version', models.PositiveBigIntegerField()),
                ('uid', models.UUIDField()),
                ('created_at', models.DateTimeField(auto_now_add=True)),
                ('updated_at', models.DateTimeField(auto_now=True)),
            ],
            options={
                'db_table': 'kluster_resources_status',
                'unique_together': {('api_group', 'kind', 'name', 'namespace', 'uid')},
            },
        ),
    ]
