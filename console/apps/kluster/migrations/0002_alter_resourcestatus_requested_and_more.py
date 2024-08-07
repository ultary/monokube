# Generated by Django 5.0.7 on 2024-07-15 01:19

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('kluster', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='resourcestatus',
            name='requested',
            field=models.JSONField(max_length=1048576),
        ),
        migrations.CreateModel(
            name='LatestRsourceKindVersion',
            fields=[
                ('resource_version', models.PositiveBigIntegerField(primary_key=True, serialize=False)),
                ('updated_at', models.DateTimeField(auto_now=True)),
            ],
            options={
                'db_table': 'kluster_latest_event_resource_version',
                'indexes': [models.Index(fields=['updated_at'], name='kluster_lat_updated_a6270f_idx')],
            },
        ),
    ]
